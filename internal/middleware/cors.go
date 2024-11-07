package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

// CORSConfig CORS 配置
type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge          int // 预检请求的缓存时间（秒）
}

// DefaultCORSConfig 默认 CORS 配置
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Authorization", 
			"Content-Type", 
			"Accept", 
			"Origin", 
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
		},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		MaxAge:          86400, // 24小时
	}
}

// CORS 创建 CORS 中间件
func CORS(config CORSConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 设置允许的源
			origin := r.Header.Get("Origin")
			if origin != "" {
				if len(config.AllowOrigins) == 0 || config.AllowOrigins[0] == "*" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
				} else {
					for _, allowedOrigin := range config.AllowOrigins {
						if origin == allowedOrigin {
							w.Header().Set("Access-Control-Allow-Origin", origin)
							break
						}
					}
				}
			}

			// 设置其他 CORS 头
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(config.AllowMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(config.AllowHeaders, ", "))
			if config.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			if len(config.ExposeHeaders) > 0 {
				w.Header().Set("Access-Control-Expose-Headers", strings.Join(config.ExposeHeaders, ", "))
			}
			if config.MaxAge > 0 {
				w.Header().Set("Access-Control-Max-Age", strconv.Itoa(config.MaxAge))
			}

			// 处理 OPTIONS 预检请求
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
