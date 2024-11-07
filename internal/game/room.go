package game

import (
	"time"
)

type RoomStatus int

const (
	RoomStatusWaiting RoomStatus = iota
	RoomStatusBanPick
	RoomStatusPlaying
	RoomStatusFinished
)

type GameRoom struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Status     RoomStatus      `json:"status"`
	GameState  *GameState      `json:"gameState"`
	Players    map[string]Team `json:"players"` // key: userId, value: playerId (white/black)
	MaxPlayers int             `json:"maxPlayers"`
	CreateTime int64           `json:"createTime"`
}

func NewGameRoom(id string, name string) *GameRoom {
	return &GameRoom{
		ID:         id,
		Name:       name,
		Status:     RoomStatusWaiting,
		Players:    make(map[string]Team),
		MaxPlayers: 2,
		CreateTime: time.Now().Unix(),
	}
}
