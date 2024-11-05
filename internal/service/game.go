package service

import (
	"fmt"
	"sync"

	"github.com/Mor1arty/WarChest/internal/game"
	"github.com/Mor1arty/WarChest/internal/game/handler"
	"github.com/Mor1arty/WarChest/pkg/utils"
)

type GameService struct {
	games    map[string]*game.GameState // key: gameId
	handlers map[game.GameActionType]handler.Handler
	mutex    sync.RWMutex
}

func NewGameService() *GameService {
	return &GameService{
		games:    make(map[string]*game.GameState),
		handlers: make(map[game.GameActionType]handler.Handler),
	}
}

// GetGame 获取指定游戏的状态
func (s *GameService) GetGame(gameID string) (*game.GameState, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	game, exists := s.games[gameID]
	return game, exists
}

// CreateNewGame 创建新游戏
func (s *GameService) CreateNewGame(players map[string]game.Team) *game.GameState {
	gamePlayers := make([]*game.Player, 0)
	for userID, team := range players {
		// 创建玩家初始单位
		supply := []*game.Unit{
			utils.GenerateUnit(game.UnitTypeLightCavalry, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeLightCavalry, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeArcher, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeArcher, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeArcher, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeArcher, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeSwordsman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeSwordsman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeSwordsman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeSwordsman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeCrossbowman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeCrossbowman, userID, game.UnitStatusInSupply),
			utils.GenerateUnit(game.UnitTypeCrossbowman, userID, game.UnitStatusInSupply),
		}

		player := game.NewPlayer(
			userID,
			userID,
			team,
			supply,
			[]*game.Unit{},
			[]*game.Unit{},
			[]*game.Unit{},
			[]*game.Unit{},
		)
		gamePlayers = append(gamePlayers, player)
	}

	// 创建新的游戏状态
	gameID := utils.GenerateUUID()
	gameState := game.NewGameState(gameID, gamePlayers, gamePlayers[0].ID)

	// 保存游戏状态
	s.mutex.Lock()
	s.games[gameID] = gameState
	s.mutex.Unlock()

	return gameState
}

// HandleAction 处理游戏动作
func (s *GameService) HandleAction(gameID string, action *game.GameAction) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	gameState, exists := s.games[gameID]
	if !exists {
		return fmt.Errorf("game not found: %s", gameID)
	}

	handler, exists := s.handlers[action.Type]
	if !exists {
		return fmt.Errorf("no handler for action type: %d", action.Type)
	}

	return handler.Handle(gameState, action.Type, action.Payload)
}
