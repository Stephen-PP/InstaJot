package api

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	db *sql.DB
}

func NewAuthMiddleware(db *sql.DB) *AuthMiddleware {
	return &AuthMiddleware{db: db}
}

func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		token := bearerToken[1]

		// TODO: Validate JWT token and extract user ID

		// For now, we'll just pass a dummy user ID
		userID := int64(1)

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
