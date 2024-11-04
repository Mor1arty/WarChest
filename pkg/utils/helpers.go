package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Mor1arty/WarChest/internal/game"
)

// GenerateUUID 随机相关
func GenerateUUID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", r.Intn(100000000))
}

func GenerateUnit(t game.UnitType, o string, s game.UnitStatus) *game.Unit {
	id := "unit_" + GenerateUUID()
	return &game.Unit{
		ID:     id,
		Type:   t,
		Owner:  o,
		Status: s,
	}
}

// ShuffleBag 棋子袋操作相关
func ShuffleBag[T any](slice []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// WrapError 错误处理相关
func WrapError(err error, message string) error {
	// 包装错误信息，添加上下文
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// 获取单位定义的辅助函数
func GetUnitDefinition(unitType game.UnitType) (game.UnitDefinition, bool) {
	def, exists := game.UnitDefinitions[unitType]
	return def, exists
}

// 获取所有单位
func GetAllUnits() []game.UnitType {
	units := make([]game.UnitType, 0, int(game.UnitTypeMax))
	for i := game.UnitType(0); i < game.UnitTypeMax; i++ {
		units = append(units, i)
	}
	return units
}

// 获取所有单位定义
func GetAllUnitDefinitions() map[game.UnitType]game.UnitDefinition {
	return game.UnitDefinitions
}
