package main

import (
	"log"
	"net/http"

	"github.com/Mor1arty/WarChest/internal/api/handler"
	"github.com/Mor1arty/WarChest/internal/middleware"
	"github.com/Mor1arty/WarChest/internal/service"
	"github.com/Mor1arty/WarChest/internal/websocket"
)

func main() {
	// 创建服务和处理器
	secretKey := []byte("12345678901234567890123456789012")
	authService := service.NewAuthService(secretKey)
	authHandler := handler.NewAuthHandler(authService)

	// 创建路由器
	mux := http.NewServeMux()

	// CORS 配置
	corsConfig := middleware.DefaultCORSConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000", // React 开发服务器
		"http://localhost:8080", // 本地开发
	}
	corsMiddleware := middleware.CORS(corsConfig)

	// 使用中间件包装每个处理器
	mux.Handle("/api/auth/validate", corsMiddleware(http.HandlerFunc(authHandler.ValidateToken)))
	mux.Handle("/api/auth/login", corsMiddleware(http.HandlerFunc(authHandler.Login)))

	// WebSocket 服务器 - 传入 authService
	ws := websocket.NewWebSocketServer(":8080", authService)
	mux.Handle("/ws", corsMiddleware(
		http.HandlerFunc(ws.HandleConnection),
	))

	// 启动服务器
	log.Printf("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
