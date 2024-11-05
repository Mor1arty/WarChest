package main

import (
	"log"
	"net/http"

	"github.com/Mor1arty/WarChest/internal/middleware"
	"github.com/Mor1arty/WarChest/internal/service"
	"github.com/Mor1arty/WarChest/internal/websocket"
)

func main() {
	// 创建认证处理器
	authService := service.NewAuthHandler()

	// 创建路由器
	mux := http.NewServeMux()

	// CORS 配置
	corsConfig := middleware.DefaultCORSConfig()

	// 可以根据需要修改配置
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",   // React 开发服务器
		"http://localhost:8080",   // 本地开发
		"https://your-domain.com", // 生产环境
	}

	// 注册路由
	mux.Handle("/api/login", middleware.CORS(corsConfig)(
		http.HandlerFunc(authService.Login),
	))

	// WebSocket 服务器 - 传入 authService
	ws := websocket.NewWebSocketServer(":8080", authService)
	mux.Handle("/ws", middleware.CORS(corsConfig)(
		http.HandlerFunc(ws.HandleConnection),
	))

	// 启动 HTTP 服务器
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
