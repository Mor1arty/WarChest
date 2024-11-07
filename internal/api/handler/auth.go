package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Mor1arty/WarChest/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GeneralResponse struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), loginReq.Username, loginReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := GeneralResponse{
		Type: "LOGIN_SUCCESS",
		Payload: struct {
			Token  string `json:"token"`
			UserID string `json:"userId"`
		}{
			Token:  token,
			UserID: loginReq.Username,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "No authorization header", http.StatusUnauthorized)
		return
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
		return
	}

	token := bearerToken[1]

	_, err := h.authService.ValidateToken(token)
	if err != nil {
		response := GeneralResponse{
			Type: "TOKEN_INVALID",
			Payload: map[string]interface{}{
				"valid": false,
				"error": err.Error(),
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	response := GeneralResponse{
		Type: "TOKEN_VALID",
		Payload: map[string]interface{}{
			"valid": true,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
