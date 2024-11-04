package service

import (
	"fmt"

	"github.com/Mor1arty/WarChest/internal/game"
	"github.com/Mor1arty/WarChest/internal/game/handler"
	"github.com/Mor1arty/WarChest/pkg/utils"
)

type GameService struct {
	gameState *game.GameState
	handlers  map[game.GameActionType]handler.Handler
}

func NewGameService() *GameService {
	gameId := utils.GenerateUUID()
	players := []*game.Player{}

	supply1 := []*game.Unit{
		utils.GenerateUnit(game.UnitTypeLightCavalry, "player1", game.UnitStatusInSupply),
		utils.GenerateUnit(game.UnitTypeArcher, "player1", game.UnitStatusInSupply),
		utils.GenerateUnit(game.UnitTypeSwordsman, "player1", game.UnitStatusInSupply),
	}

	supply2 := []*game.Unit{
		utils.GenerateUnit(game.UnitTypeLightCavalry, "player2", game.UnitStatusInSupply),
		utils.GenerateUnit(game.UnitTypeArcher, "player2", game.UnitStatusInSupply),
		utils.GenerateUnit(game.UnitTypeSwordsman, "player2", game.UnitStatusInSupply),
	}
	players = append(players,
		game.NewPlayer("player1", "player1", game.TeamWhite, supply1, []*game.Unit{}, []*game.Unit{}, []*game.Unit{}, []*game.Unit{}),
		game.NewPlayer("player2", "player2", game.TeamBlack, supply2, []*game.Unit{}, []*game.Unit{}, []*game.Unit{}, []*game.Unit{}),
	)

	service := &GameService{
		gameState: game.NewGameState(gameId, players, "player1"),
		handlers:  make(map[game.GameActionType]handler.Handler),
	}

	// 注册处理器
	service.registerHandlers()
	return service
}

func (s *GameService) GetGameState() *game.GameState {
	return s.gameState
}

func (s *GameService) registerHandlers() {
	s.handlers[game.GameActionTypeClear] = handler.HandlerFunc(handler.ClearGameHandler)
	// ... 注册其他处理器
}

func (s *GameService) HandleAction(action *game.GameAction) error {
	handler, exists := s.handlers[action.Type]
	if !exists {
		return fmt.Errorf("no handler for action type: %d", action.Type)
	}

	return handler.Handle(action.Type, action.Payload)
}
