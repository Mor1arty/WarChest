package handler

import "github.com/Mor1arty/WarChest/internal/game"

// Handler 游戏动作处理器接口
type Handler interface {
	Handle(gameState *game.GameState, actionType game.GameActionType, payload interface{}) error
}

// HandlerFunc 处理器函数类型
type HandlerFunc func(gameState *game.GameState, actionType game.GameActionType, payload interface{}) error

// Handle 实现 Handler 接口
func (f HandlerFunc) Handle(gameState *game.GameState, actionType game.GameActionType, payload interface{}) error {
	return f(gameState, actionType, payload)
}
