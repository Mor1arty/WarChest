package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateUUID 随机相关
func GenerateUUID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", r.Intn(100000000))
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
