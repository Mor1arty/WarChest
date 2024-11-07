package game

import (
	"sync"
)

type GameState struct {
	mutex sync.RWMutex `json:"-"` // 互斥锁，不进行JSON序列化

	// 游戏基础信息
	GameID string `json:"gameId"`

	// 回合信息
	CurrentTurn   int    `json:"currentTurn"`
	CurrentPlayer string `json:"currentPlayer"`
	Initiative    string `json:"initiative"`

	// 玩家信息
	Players []*Player `json:"players"`

	// 单位信息
	Units map[string]*Unit `json:"units"` // key: unitId

	// 棋盘信息
	Board *Board `json:"board"`

	// 游戏历史记录
	History []HistoryRecord `json:"history"`
}

// NewGameState 创建新的游戏状态
func NewGameState(gameID string, players []*Player, initiative string) *GameState {
	units := make(map[string]*Unit)
	for _, player := range players {
		for _, unit := range player.Supply {
			units[unit.ID] = unit
		}
		for _, unit := range player.Hand {
			units[unit.ID] = unit
		}
		for _, unit := range player.Bag {
			units[unit.ID] = unit
		}
		for _, unit := range player.DiscardPile {
			units[unit.ID] = unit
		}
	}
	return &GameState{
		GameID:      gameID,
		CurrentTurn: 0,
		Initiative:  initiative,
		Players:     players,
		Units:       make(map[string]*Unit),
		Board:       CreateBoard(len(players)),
		History:     make([]HistoryRecord, 0),
	}
}

// 游戏状态的方法

// GetPlayer 获取玩家信息
func (gs *GameState) GetPlayer(playerID string) (*Player, bool) {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()
	for _, player := range gs.Players {
		if player.ID == playerID {
			return player, true
		}
	}
	return nil, false
}

// GetUnit 获取单位信息
func (gs *GameState) GetUnit(unitID string) (*Unit, bool) {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()
	unit, exists := gs.Units[unitID]
	return unit, exists
}

// GetCell 获取指定位置的格子信息
func (gs *GameState) GetCell(position string) (*BoardCell, bool) {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()
	if cell, exists := gs.Board.Cells[position]; exists {
		return cell, true
	}
	return nil, false
}

// AddHistoryRecord 添加历史记录
func (gs *GameState) AddHistoryRecord(actions []GameAction) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
}

// NextTurn 进入下一回合
func (gs *GameState) NextTurn() {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	initiative := gs.Initiative
	initiativeIndex := 0
	for i, player := range gs.Players {
		if player.ID == initiative {
			initiativeIndex = i
			break
		}
	}
	gs.CurrentTurn++
	gs.CurrentPlayer = gs.Players[(initiativeIndex+gs.CurrentTurn)%len(gs.Players)].ID
}

// AddUnit 添加单位
func (gs *GameState) AddUnit(unit *Unit) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	gs.Units[unit.ID] = unit
}

// RemoveUnit 移除单位
func (gs *GameState) RemoveUnit(unitID string) {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	delete(gs.Units, unitID)
}

// SetCellControl 设置格子控制权
func (gs *GameState) SetCellControl(position string, playerID string) bool {
	gs.mutex.Lock()
	defer gs.mutex.Unlock()
	if cell, exists := gs.GetCell(position); exists {
		cell.ControlledBy = &playerID
		gs.Board.Cells[position] = cell
		return true
	}
	return false
}

// GetState 获取完整的游戏状态（用于序列化发送给前端）
func (gs *GameState) GetState() *GameState {
	gs.mutex.RLock()
	defer gs.mutex.RUnlock()
	return gs
}
