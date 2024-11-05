package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/net/context"
)

const secretString = "12345678901234567890123456789012"

// LoginRequest 定义登录请求的结构
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 定义登录响应的结构
type LoginResponse struct {
	Type    string `json:"type"`
	Payload struct {
		Token  string `json:"token"`
		UserID string `json:"userId"`
	} `json:"payload"`
}

// AuthHandler 处理认证相关的请求
type AuthHandler struct {
	jwtSecret []byte
}

// NewAuthHandler 创建新的认证处理器
func NewAuthHandler() *AuthHandler {
	// 从环境变量获取 secret
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		// 如果环境变量未设置，使用默认值（仅用于开发）
		secret = []byte(secretString)
	}
	return &AuthHandler{
		jwtSecret: secret,
	}
}

func (h *AuthHandler) validateCredentials(username, password string) bool {
	// 这里应该替换为实际的用户验证逻辑
	user := map[string]string{
		"admin": "123",
		"user1": "123",
		"user2": "123",
	}
	return user[username] == password
}

// Login 处理登录请求
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// 只允许 POST 方法
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 验证用户名和密码
	if !h.validateCredentials(loginReq.Username, loginReq.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 生成用户ID（这里使用用户名作为示例）
	userID := loginReq.Username

	// 生成 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,                                // 用户ID
		"exp": time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
		"iat": time.Now().Unix(),                     // 签发时间
	})

	// 签名 token
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		log.Printf("生成 token 失败: %v", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// 返回登录响应
	loginResp := LoginResponse{
		Type: "LOGIN_SUCCESS",
		Payload: struct {
			Token  string `json:"token"`
			UserID string `json:"userId"`
		}{
			Token:  tokenString,
			UserID: userID,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResp)
}

// ValidateToken 验证 token
func (h *AuthHandler) ValidateToken(tokenString string) (*jwt.Token, error) {
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

func (h *AuthHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 获取 token
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// 2. 验证 token
		token, err := h.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 3. 获取用户信息
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		role := claims["role"].(string)

		// 4. 将用户信息添加到请求上下文
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)

		// 5. 继续处理请求
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
