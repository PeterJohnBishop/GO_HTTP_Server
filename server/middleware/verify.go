package middleware

import (
	"context"
	"encoding/json"
	"free-adventure-go/main.go/auth"
	"net/http"
	"strings"
)

//////////////////////////////////////// IN PROGRESS ////////////////////////////////////////

type contextKey string

const userClaimsKey contextKey = "userClaims"

// authenticate JWT token
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
		userClaims := auth.ParseAccessToken(token)
		if userClaims == nil {
			http.Error(w, `{"error": "Failed to verify token!"}`, http.StatusUnauthorized)
			return
		}

		// Store claims in context
		ctx := context.WithValue(r.Context(), userClaimsKey, userClaims)
		r = r.WithContext(ctx) // Assign the modified request back

		next.ServeHTTP(w, r)
	})
}

// authenticate Refresh Token
type VerifyRefreshRequest struct {
	Token string `json:"token"`
}

func VerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req VerifyRefreshRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		token := req.Token
		claims := auth.ParseRefreshToken(token)
		if claims == nil {
			http.Error(w, `{"error": "Failed to verify token!"}`, http.StatusUnauthorized)
			return
		}
		// Add userClaims to the request context
		next.ServeHTTP(w, r)
	})
}

// generate new access token

// generate new refresh token
