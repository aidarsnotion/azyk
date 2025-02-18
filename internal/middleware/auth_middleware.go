package middleware

import (
	"context"
	"net/http"
	"strings"

	"azyk/util"
)

type contextKey string

const (
	ContextUserID = contextKey("user_id")
	ContextRole   = contextKey("role")
)

// Middleware для проверки JWT токена
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Ожидаемый формат: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims, err := util.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем user_id и роль в контекст
		ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
		ctx = context.WithValue(ctx, ContextRole, claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
