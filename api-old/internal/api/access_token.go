package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	util "github.com/stephen-pp/instajot/internal/api/util"
	"github.com/stephen-pp/instajot/internal/crypto"
)

func (api *APIServer) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.WriteFailureResponse(
			w, http.StatusMethodNotAllowed,
			util.NewErrorResponse("Method not allowed", "METHOD_NOT_ALLOWED", ""),
		)
		return
	}

	var request struct {
		Token string `json:"token"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		util.WriteFailureResponse(
			w, http.StatusBadRequest,
			util.NewErrorResponse("Invalid request body", "INVALID_REQUEST_BODY", err.Error()),
		)
		return
	}

	row := api.DB.QueryRow("SELECT user_id, expires_at FROM access_tokens WHERE token = ?", request.Token)
	var userID int64
	var expiresAt time.Time
	err = row.Scan(&userID, &expiresAt)
	if err != nil {
		util.WriteFailureResponse(
			w, http.StatusUnauthorized,
			util.NewErrorResponse("Invalid token", "INVALID_TOKEN", err.Error()),
		)
		return
	}

	if time.Now().After(expiresAt) {
		util.WriteFailureResponse(
			w, http.StatusUnauthorized,
			util.NewErrorResponse("Token expired", "TOKEN_EXPIRED", ""),
		)
		return
	}

	newToken, err := crypto.GenerateRandomString(20)
	if err != nil {
		log.Printf("Error generating new token: %v", err)
		util.WriteFailureResponse(
			w, http.StatusInternalServerError,
			util.NewErrorResponse("Internal server error", "INTERNAL_SERVER_ERROR", err.Error()),
		)
		return
	}
	newExpiresAt := time.Now().Add(48 * time.Hour)

	_, err = api.DB.Exec("UPDATE access_tokens SET token = ?, expires_at = ? WHERE token = ?", newToken, newExpiresAt, request.Token)
	if err != nil {
		log.Printf("Error updating token: %v", err)
		util.WriteFailureResponse(
			w, http.StatusInternalServerError,
			util.NewErrorResponse("Internal server error", "INTERNAL_SERVER_ERROR", err.Error()),
		)
		return
	}

	response := struct {
		Token     string `json:"token"`
		ExpiresAt string `json:"expires_at"`
	}{
		Token:     newToken,
		ExpiresAt: newExpiresAt.Format(time.RFC3339),
	}

	util.WriteSuccessResponse(w, response)
}
