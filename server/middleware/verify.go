package middleware

import (
	"free-adventure-go/main.go/auth"
	"net/http"
	"strings"
)

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix((r.URL.Path), "/register") || strings.HasPrefix((r.URL.Path), "/login") {
			next.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Authentication Header is missing!"}`, http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		err := auth.VerifyToken(token)
		if err != nil {
			http.Error(w, `{"error": "Failed to verify token!"}`, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
