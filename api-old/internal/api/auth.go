package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	util "github.com/stephen-pp/instajot/internal/api/util"
	"github.com/stephen-pp/instajot/internal/crypto"
	"github.com/stephen-pp/instajot/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (api *APIServer) registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.WriteFailureResponse(w, http.StatusMethodNotAllowed, util.NewErrorResponse("Method not allowed", "METHOD_NOT_ALLOWED", ""))
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.WriteFailureResponse(w, http.StatusBadRequest, util.NewErrorResponse("Invalid request body", "INVALID_REQUEST_BODY", ""))
		return
	}
	// Validate user input
	if user.Username == "" || user.Email == "" || user.Password == "" {
		util.WriteFailureResponse(w, http.StatusBadRequest, util.NewErrorResponse("Username, email, and password are required", "VALIDATION_ERROR", ""))
		return
	}
	// Hash the password
	hashedRaw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Internal server error", "INTERNAL_SERVER_ERROR", ""))
		return
	}
	hashedPassword := string(hashedRaw)
	// Create user in the database
	result, err := api.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		user.Username, user.Email, hashedPassword)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Error creating user", "DATABASE_ERROR", ""))
		return
	}
	userID, _ := result.LastInsertId()
	// Generate access token
	token, err := crypto.GenerateRandomString(20) // Implement this function to generate a unique token
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Error generating access token", "TOKEN_GENERATION_ERROR", ""))
		return
	}
	expiresAt := time.Now().Add(48 * time.Hour) // Token expires in 48 hours
	// Store access token in the database
	_, err = api.DB.Exec("INSERT INTO access_tokens (user_id, token, expires_at) VALUES (?, ?, ?)",
		userID, token, expiresAt)
	if err != nil {
		log.Printf("Error storing access token: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Error storing access token", "DATABASE_ERROR", ""))
		return
	}
	// Prepare response
	response := struct {
		UserID    int64  `json:"user_id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Token     string `json:"token"`
		ExpiresAt string `json:"expires_at"`
	}{
		UserID:    userID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		ExpiresAt: expiresAt.Format(time.RFC3339),
	}
	util.WriteSuccessResponse(w, response)
}

func (api *APIServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.WriteFailureResponse(w, http.StatusMethodNotAllowed, util.NewErrorResponse("Method not allowed", "METHOD_NOT_ALLOWED", ""))
		return
	}
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.WriteFailureResponse(w, http.StatusBadRequest, util.NewErrorResponse("Invalid request body", "INVALID_REQUEST_BODY", ""))
		return
	}
	// Validate user input
	if user.Username == "" || user.Password == "" {
		util.WriteFailureResponse(w, http.StatusBadRequest, util.NewErrorResponse("Username and password are required", "VALIDATION_ERROR", ""))
		return
	}
	// Retrieve user from the database
	row := api.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", user.Username)
	var dbUser models.User
	err = row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Email, &dbUser.Password)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Invalid username or password", "INVALID_CREDENTIALS", ""))
		return
	}
	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		log.Printf("Error comparing passwords: %v", err)
		util.WriteFailureResponse(w, http.StatusUnauthorized, util.NewErrorResponse("Invalid username or password", "INVALID_CREDENTIALS", ""))
		return
	}
	// Generate access token
	token, err := crypto.GenerateRandomString(20) // Implement this function to generate a unique token
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Error generating access token", "TOKEN_GENERATION_ERROR", ""))
		return
	}
	expiresAt := time.Now().Add(48 * time.Hour) // Token expires in 48 hours
	// Store access token in the database
	_, err = api.DB.Exec("INSERT INTO access_tokens (user_id, token, expires_at) VALUES (?, ?, ?)",
		dbUser.ID, token, expiresAt)
	if err != nil {
		log.Printf("Error storing access token: %v", err)
		util.WriteFailureResponse(w, http.StatusInternalServerError, util.NewErrorResponse("Error storing access token", "DATABASE_ERROR", ""))
		return
	}
	// Prepare response
	response := struct {
		UserID    int64  `json:"user_id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Token     string `json:"token"`
		ExpiresAt string `json:"expires_at"`
	}{
		UserID:    dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Token:     token,
		ExpiresAt: expiresAt.Format(time.RFC3339),
	}

	util.WriteSuccessResponse(w, response)
}
