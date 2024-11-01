package main

import (
	"log"

	"github.com/Mor1arty/WarChest/internal/websocket"
)

func main() {
	ws := websocket.NewWebSocketServer(":8080")
	if err := ws.Start(); err != nil {
		log.Fatalf("WebSocket 服务器启动失败: %v", err)
	}
}
