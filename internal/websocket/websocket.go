package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Mor1arty/WarChest/internal/service"
	"github.com/Mor1arty/WarChest/pkg/utils"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type WebSocketServer struct {
	upgrader    websocket.Upgrader
	port        string
	clients     map[string]*Client
	mutex       sync.RWMutex
	gameService *service.GameService
}

func checkOrigin(r *http.Request) bool {
	return true
}

// NewWebSocketServer 创建一个新的 WebSocket 服务器实例
func NewWebSocketServer(port string) *WebSocketServer {
	return &WebSocketServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: checkOrigin,
		},
		port:        port,
		clients:     make(map[string]*Client),
		gameService: service.NewGameService(),
	}
}

// handleConnection 处理单个 WebSocket 连接
func (ws *WebSocketServer) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("升级连接失败: %v", err)
		return
	}

	// 为新客户端创建唯一ID
	clientID := utils.GenerateUUID()
	client := &Client{
		ID:   clientID,
		Conn: conn,
	}

	// 注册客户端
	ws.registerClient(client)
	defer ws.unregisterClient(client)

	// 发送欢迎消息
	welcomeData := map[string]interface{}{
		"type": "WELCOME",
		"payload": map[string]interface{}{
			"success":  true,
			"clientID": clientID,
		},
	}
	jsonWelcomeData, err := json.Marshal(welcomeData)
	if err != nil {
		log.Printf("JSON序列化错误: %v", err)
		return
	}
	ws.SendToClient(client, string(jsonWelcomeData))

	// 处理接收到的消息
	ws.handleMessages(client)
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

// 修改 handleMessages 函数
func (ws *WebSocketServer) handleMessages(client *Client) {
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息错误: %v", err)
			break
		}

		// 解析客户端消息
		var clientMsg ClientMessage
		if err := json.Unmarshal(message, &clientMsg); err != nil {
			log.Printf("解析消息失败: %v", err)
			continue
		}

		// 根据消息类型处理
		switch clientMsg.Type {
		case "UPDATE_GAME_STATE":
			fmt.Println("更新游戏状态")
			gameState := ws.gameService.GetGameState()
			response := map[string]interface{}{
				"type": "UPDATE_GAME_STATE",
				"payload": map[string]interface{}{
					"success":   true,
					"changes":   []string{},
					"gameState": gameState,
				},
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				log.Printf("JSON序列化错误: %v", err)
				continue
			}

			ws.SendToClient(client, string(jsonResponse))
		default:
			log.Printf("未知的消息类型: %s", clientMsg.Type)
		}
	}
}

// SendToClient 向特定客户端发送消息
func (ws *WebSocketServer) SendToClient(client *Client, message string) error {
	return client.Conn.WriteMessage(websocket.TextMessage, []byte(message))
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

// Start 启动 WebSocket 服务器
func (ws *WebSocketServer) Start() error {
	http.HandleFunc("/ws", ws.handleConnection)
	log.Printf("WebSocket 服务器启动在 %s 端口...", ws.port)
	return http.ListenAndServe(ws.port, nil)
}

// GetConnectedClients 获取当前连接的客户端数量
func (ws *WebSocketServer) GetConnectedClients() int {
	ws.mutex.RLock()
	defer ws.mutex.RUnlock()
	return len(ws.clients)
}
