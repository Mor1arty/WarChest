package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	jwtSecret []byte
}

func NewAuthService(jwtSecret []byte) *AuthService {
	return &AuthService{
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	if !s.validateCredentials(username, password) {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) validateCredentials(username, password string) bool {
	// 验证逻辑
	return true
}

// ValidateToken 验证 token
func (h *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return h.jwtSecret, nil
	})

	if err != nil {
		log.Printf("验证 token 失败 %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 检查是否过期
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token expired")
			}
		}

		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}
