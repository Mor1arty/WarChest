package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Mor1arty/WarChest/internal/game"
	"github.com/Mor1arty/WarChest/internal/service"
	"github.com/Mor1arty/WarChest/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

const (
	pingInterval = 30 * time.Second // 发送 ping 的间隔
	pongWait     = 60 * time.Second // 等待 pong 响应的超时时间
	writeWait    = 10 * time.Second // 写入超时
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	UserID string
	Send   chan []byte // 消息发送通道
}

type AuthPayload struct {
	Token string `json:"token"`
}

type WebSocketServer struct {
	upgrader    websocket.Upgrader
	port        string
	clients     map[string]*Client
	mutex       sync.RWMutex
	gameService *service.GameService
	authService *service.AuthService
	rooms       map[string]*game.GameRoom // 游戏房间映射
	userRooms   map[string]string         // 用户ID到房间ID的映射
	roomMutex   sync.RWMutex
}

func checkOrigin(r *http.Request) bool {
	return true
}

// NewWebSocketServer 创建一个新的 WebSocket 服务器实例
func NewWebSocketServer(port string, authService *service.AuthService) *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: checkOrigin,
		},
		port:        port,
		clients:     make(map[string]*Client),
		gameService: service.NewGameService(),
		authService: authService,
		rooms:       make(map[string]*game.GameRoom),
		userRooms:   make(map[string]string),
		roomMutex:   sync.RWMutex{},
	}
}

// HandleConnection 处理新的 WebSocket 连接
func (ws *WebSocketServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// 1. 从 URL 参数获取 token
	token := r.URL.Query().Get("token")
	if token == "" {
		log.Printf("未提供 token")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 2. 使用 authService 验证 token
	jwtToken, err := ws.authService.ValidateToken(token)
	if err != nil {
		log.Printf("token 验证失败: %v", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// 3. 获取用户信息
	claims := jwtToken.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)

	// 4. 升级连接
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}

	client := &Client{
		ID:     utils.GenerateUUID(),
		Conn:   conn,
		UserID: userID,
		Send:   make(chan []byte, 256),
	}

	// 设置读取超时
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// 注册客户端
	ws.registerClient(client)

	// 启动读写 goroutines
	go ws.writePump(client)
	go ws.readPump(client)

	// 7. 发送连接成功消息
	ws.SendToClient(client, `{"type":"CONNECTED","payload":{"message":"连接成功","userId":"`+userID+`"}}`)
}

// registerClient 注册新客户端
func (ws *WebSocketServer) registerClient(client *Client) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	ws.clients[client.ID] = client
	log.Printf("新客户端连接: %s", client.ID)
}

// unregisterClient 注销客户端
func (ws *WebSocketServer) unregisterClient(client *Client) {
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	if _, ok := ws.clients[client.ID]; ok {
		client.Conn.Close()
		delete(ws.clients, client.ID)
		log.Printf("客户端断开连接: %s", client.ID)
	}
}

// 定义客户端消息结构
type ClientMessage struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// readPump 处理从客户端读取消息
func (ws *WebSocketServer) readPump(client *Client) {
	defer func() {
		ws.unregisterClient(client)
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("读取错误: %v", err)
			}
			break
		}

		// 处理消息
		var clientMsg ClientMessage
		if err := json.Unmarshal(message, &clientMsg); err != nil {
			log.Printf("解析消息失败: %v", err)
			continue
		}

		// 处理消息
		switch clientMsg.Type {
		case "CREATE_ROOM":
			log.Printf(client.UserID + " Create room")
			ws.handleCreateRoom(client)
		case "JOIN_ROOM":
			log.Printf(client.UserID + " Join room")
			if ws.checkGameMatch(client) {
				continue
			} else {
				ws.handleGameMatch(client)
			}
		case "LEAVE_ROOM":
			log.Printf(client.UserID + " Leave room")
			ws.handleLeaveRoom(client)
		case "UPDATE_GAME_STATE":
			ws.roomMutex.RLock()
			roomID, exists := ws.userRooms[client.UserID]
			if !exists {
				ws.roomMutex.RUnlock()
				continue
			}
			room, exists := ws.rooms[roomID]
			if !exists || room.GameState == nil {
				ws.roomMutex.RUnlock()
				continue
			}
			gameID := room.GameState.GameID
			ws.roomMutex.RUnlock()

			if gameState, exists := ws.gameService.GetGame(gameID); exists {
				response := map[string]interface{}{
					"type": "UPDATE_GAME_STATE",
					"payload": map[string]interface{}{
						"success":   true,
						"changes":   []string{},
						"gameState": gameState,
					},
				}
				jsonResponse, _ := json.Marshal(response)
				client.Send <- jsonResponse
			}
		case "PING":
			response := ClientMessage{
				Type: "PONG",
				Payload: map[string]interface{}{
					"time": time.Now().Unix(),
				},
			}
			jsonResponse, _ := json.Marshal(response)
			client.Send <- jsonResponse
		}
	}
}

// writePump 处理向客户端发送消息
func (ws *WebSocketServer) writePump(client *Client) {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Send channel 已关闭
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SendToClient 发送消息给指定客户端
func (ws *WebSocketServer) SendToClient(client *Client, message string) error {
	select {
	case client.Send <- []byte(message):
		return nil
	default:
		ws.unregisterClient(client)
		return fmt.Errorf("client message buffer is full")
	}
}

// BroadcastMessage 向所有客户端广播消息
func (ws *WebSocketServer) BroadcastMessage(messageType int, message []byte) {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()

	for _, client := range ws.clients {
		err := client.Conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("发送消息到客户端 %s 失败: %v", client.ID, err)
		}
	}
}

// 检查是否有正在进行中的游戏
func (ws *WebSocketServer) checkGameMatch(client *Client) bool {
	ws.roomMutex.Lock()
	defer ws.roomMutex.Unlock()

	if roomID, exists := ws.userRooms[client.UserID]; exists {
		if room, exists := ws.rooms[roomID]; exists {
			log.Printf("room found: " + room.ID)
			if room.Status == game.RoomStatusPlaying {
				gameState, _ := json.Marshal(room.GameState)
				log.Printf("send game state to client ")
				ws.SendToClient(client, `{"type":"UPDATE_GAME_STATE", "payload": {"roomId": "`+room.ID+`", "gameState": `+string(gameState)+`}}`)
				return true
			}
		}
	}
	return false
}

// 创建房间
func (ws *WebSocketServer) createRoom(client *Client) string {
	roomID := utils.GenerateUUID()
	availableRoom := game.NewGameRoom(roomID, "房间-"+roomID)
	ws.rooms[roomID] = availableRoom
	ws.userRooms[client.UserID] = roomID
	ws.SendToClient(client, `{"type":"ROOM_CREATED", "payload": {"roomId": "`+roomID+`"}}`)
	return roomID
}

// 将玩家加入指定房间
func (ws *WebSocketServer) addUserToRoom(client *Client, room *game.GameRoom) bool {
	if len(room.Players) >= room.MaxPlayers {
		return false
	}

	playerRole := game.TeamWhite
	if len(room.Players) > 0 {
		for _, role := range room.Players {
			if role == game.TeamBlack {
				playerRole = game.TeamWhite
			} else {
				playerRole = game.TeamBlack
			}
		}
	} else {
		if rand.Intn(2) == 0 {
			playerRole = game.TeamBlack
		}
	}
	room.Players[client.UserID] = playerRole
	ws.userRooms[client.UserID] = room.ID

	ws.SendToClient(client, `{"type":"JOIN_ROOM", "payload": {"success": true, "roomId": "`+room.ID+`"}}`)
	return true
}

// 处理创建游戏房间
func (ws *WebSocketServer) handleCreateRoom(client *Client) {
	roomID := ws.createRoom(client)
	ws.addUserToRoom(client, ws.rooms[roomID])
}

// 处理游戏匹配
func (ws *WebSocketServer) handleGameMatch(client *Client) {
	ws.roomMutex.Lock()
	defer ws.roomMutex.Unlock()

	// 1. 查找等待中的房间
	var availableRoom *game.GameRoom
	for _, room := range ws.rooms {
		if room.Status == game.RoomStatusWaiting && len(room.Players) < room.MaxPlayers {
			availableRoom = room
			break
		}
	}

	// 2. 如果没有可用房间，创建新房间
	if availableRoom == nil {
		roomID := ws.createRoom(client)
		availableRoom = ws.rooms[roomID]
	}

	success := ws.addUserToRoom(client, availableRoom)
	if success {
		ws.SendToClient(client, `{"type":"JOIN_ROOM", "payload": {"success": true, "roomId": "`+availableRoom.ID+`"}}`)
	} else {
		ws.SendToClient(client, `{"type":"JOIN_ROOM", "payload": {"success": false, "roomId": "`+availableRoom.ID+`"}}`)
	}

	// 4. 如果房间满员，开始游戏
	if len(availableRoom.Players) == availableRoom.MaxPlayers {
		log.Printf("create new game")
		availableRoom.Status = game.RoomStatusPlaying
		availableRoom.GameState = ws.gameService.CreateNewGame(availableRoom.Players)

		// 通知房间内所有玩家游戏开始
		ws.broadcastToRoom(availableRoom.ID, map[string]interface{}{
			"type": "UPDATE_GAME_STATE",
			"payload": map[string]interface{}{
				"roomId":    availableRoom.ID,
				"gameState": availableRoom.GameState,
			},
		})
	} else {
		// 通知玩家等待对手
		log.Printf("waiting for opponent: " + availableRoom.ID)
		ws.SendToClient(client, `{"type":"WAITING_FOR_OPPONENT","payload":{"roomId":"`+availableRoom.ID+`"}}`)
	}
}

// 广播消息给房间内的所有玩家
func (ws *WebSocketServer) broadcastToRoom(roomID string, message interface{}) {
	room, exists := ws.rooms[roomID]

	if !exists {
		log.Printf("room not found: " + roomID)
		return
	}
	log.Printf("broadcast to room: " + roomID)

	jsonMessage, _ := json.Marshal(message)

	for userID := range room.Players {
		for _, client := range ws.clients {
			if client.UserID == userID {
				client.Send <- jsonMessage
				break
			}
		}
	}
}

// 离开房间
func (ws *WebSocketServer) handleLeaveRoom(client *Client) {
	ws.roomMutex.Lock()
	defer ws.roomMutex.Unlock()
	// 清理房间关系
	if roomID, exists := ws.userRooms[client.UserID]; exists {
		if room, exists := ws.rooms[roomID]; exists {
			// 从房间中移除玩家
			delete(room.Players, client.UserID)

			// 如果房间空了，删除房间
			if len(room.Players) == 0 {
				delete(ws.rooms, roomID)
			} else {
				// 重制房间状态
				room.Status = game.RoomStatusWaiting
				if room.GameState != nil {
					ws.gameService.RemoveGame(room.GameState.GameID)
					room.GameState = nil
				}
				ws.broadcastToRoom(roomID, map[string]interface{}{
					"type": "UPDATE_ROOM_STATE",
					"payload": map[string]interface{}{
						"roomId":     roomID,
						"roomStatus": room.Status,
					},
				})
			}
		}
		// 删除用户到房间的映射
		delete(ws.userRooms, client.UserID)
	}
}

func (ws *WebSocketServer) cleanupRooms() {
	// ws.roomMutex.Lock()
	// defer ws.roomMutex.Unlock()

	// for roomID, room := range ws.rooms {
	// 	if room.Status == game.RoomStatusFinished {
	// 		// 清理游戏状态
	// 		if room.GameState != nil {
	// 			ws.gameService.RemoveGame(room.GameState.GameID)
	// 		}
	// 		// 清理用户到房间的映射
	// 		for userID := range room.Players {
	// 			delete(ws.userRooms, userID)
	// 		}
	// 		// 删除房间
	// 		delete(ws.rooms, roomID)
	// 	}
	// }
}
