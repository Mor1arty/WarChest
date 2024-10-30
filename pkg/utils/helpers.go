package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// 坐标相关
type Position struct {
	X int
	Y int
}

// 检查坐标是否在棋盘范围内
func IsPositionValid(pos Position, boardSize int) bool {
	return pos.X >= 0 && pos.X < boardSize && pos.Y >= 0 && pos.Y < boardSize
}

// 计算两点间的距离
func CalculateDistance(pos1, pos2 Position) int {
	// 使用曼哈顿距离或其他适合的距离计算方法
	dx := abs(pos1.X - pos2.X)
	dy := abs(pos1.Y - pos2.Y)
	return dx + dy
}

// 随机相关
func GenerateRandomID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", r.Intn(100000000))
}

// 棋子袋操作相关
func ShuffleBag[T any](slice []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// 错误处理相关
func WrapError(err error, message string) error {
	// 包装错误信息，添加上下文
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// 私有辅助函数
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
