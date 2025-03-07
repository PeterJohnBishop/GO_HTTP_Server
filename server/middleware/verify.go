package middleware

import (
	"context"
	"encoding/json"
	"free-adventure-go/main.go/auth"
	"net/http"
	"strings"
)

//////////////////////////////////////// IN PROGRESS ////////////////////////////////////////

// authenticate JWT token
func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		userClaims := auth.ParseAccessToken(token)
		if userClaims == nil {
			http.Error(w, `{"error": "Failed to verify token!"}`, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// authenticate Refresh Token
type VerifyRefreshRequest struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type ContextKey string

const userIDKey ContextKey = "userID"

func VerifyRefreshToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req VerifyRefreshRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		token := req.Token
		id := req.ID

		claims := auth.ParseRefreshToken(token)
		if claims == nil {
			http.Error(w, `{"error": "Failed to verify token!"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, id)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
