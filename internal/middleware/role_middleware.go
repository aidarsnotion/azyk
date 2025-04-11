package middleware

import (
	"net/http"
)

// Проверка доступа по ролям
func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(ContextRole).(string)
			if !ok {
				http.Error(w, "Forbidden: no role found", http.StatusForbidden)
				return
			}

			// Проверяем, входит ли роль пользователя в список разрешенных
			for _, requiredRole := range requiredRoles {
				if role == requiredRole {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
		})
	}
}
