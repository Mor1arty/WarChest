package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/Mor1arty/WarChest/pkg/utils"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type WebSocketServer struct {
	upgrader websocket.Upgrader
	port     string
	clients  map[string]*Client
	mutex    sync.RWMutex
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
		port:     port,
		clients:  make(map[string]*Client),
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
	clientID := utils.GenerateRandomID()
	client := &Client{
		ID:   clientID,
		Conn: conn,
	}

	// 注册客户端
	ws.registerClient(client)
	defer ws.unregisterClient(client)

	// 发送欢迎消息
	ws.SendToClient(client, "欢迎连接到服务器！你的ID是: "+clientID)

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

// handleMessages 处理来自客户端的消息
func (ws *WebSocketServer) handleMessages(client *Client) {
	for {
		messageType, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息错误: %v", err)
			break
		}

		// 处理接收到的消息
		log.Printf("收到来自 %s 的消息: %s", client.ID, string(message))

		// 广播消息给所有客户端
		ws.BroadcastMessage(messageType, message)
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