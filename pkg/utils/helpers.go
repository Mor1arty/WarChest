package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	gamemap "github.com/Mor1arty/WarChest/internal/game"
)

// GenerateRandomID 随机相关
func GenerateRandomID() string {
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

func Cube2String(c gamemap.Position) string {
	return strings.Join([]string{string(rune(c.Q)), string(rune(c.R)), string(rune(c.S))}, ",")
}

func String2Cube(s string) gamemap.Position {
	cube := strings.Split(s, ",")
	_q, qerr := strconv.Atoi(cube[0])
	_r, rerr := strconv.Atoi(cube[1])
	_s, serr := strconv.Atoi(cube[2])
	if qerr != nil || rerr != nil || serr != nil {
		return gamemap.Position{}
	}

	return gamemap.Position{
		Q: _q,
		R: _r,
		S: _s,
	}
}
