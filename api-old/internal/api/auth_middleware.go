package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	util "github.com/stephen-pp/instajot/internal/api/util"
	"github.com/stephen-pp/instajot/internal/models"
)

func (api *APIServer) authenticateAccessToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Missing authorization header", "unauthorized", ""))
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Invalid authorization header", "unauthorized", ""))
			return
		}

		token := bearerToken[1]

		var userID int64
		var expiresAt time.Time
		err := api.DB.QueryRow("SELECT user_id, expires_at FROM access_tokens WHERE token = ?", token).Scan(&userID, &expiresAt)
		if err != nil {
			util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Invalid token", "unauthorized", ""))
			return
		}

		if time.Now().After(expiresAt) {
			util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Token expired", "unauthorized", ""))
			return
		}

		// Pull the entire user from the database
		var user models.User
		err = api.DB.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Failed to fetch user", "internal_error", ""))
			return
		}

		// Set the user model in the context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
