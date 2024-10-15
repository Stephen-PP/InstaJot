package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthHandlers(t *testing.T) {
	db := setupTests()

	hashedRaw, err := bcrypt.GenerateFromPassword([]byte("validpassword"), 14)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ('validuser', 'validemail', ?)", hashedRaw)
	if err != nil {
		t.Fatalf("failed to insert user: %v", err)
	}

	t.Run("MethodNotAllowed", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:9183/auth/login", nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})

	t.Run("InvalidRequestBody", func(t *testing.T) {
		body := []byte("invalid")
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/login", bytes.NewBuffer(body))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("InvalidCredentialsUsername", func(t *testing.T) {
		body := map[string]string{"username": "invalid", "password": "invalid"}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/login", bytes.NewBuffer(bodyBytes))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
		}
	})

	// Test succesful username wrong password
	t.Run("InvalidCredentialsPassword", func(t *testing.T) {
		body := map[string]string{"username": "validuser", "password": "invalid"}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/login", bytes.NewBuffer(bodyBytes))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to send request: %v", err)
		}

		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
		}
	})

	t.Run("SuccessfulLogin", func(t *testing.T) {
		body := map[string]string{"username": "validuser", "password": "validpassword"}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/login", bytes.NewBuffer(bodyBytes))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response struct {
			Data struct {
				UserID    int64  `json:"user_id"`
				Username  string `json:"username"`
				Email     string `json:"email"`
				Token     string `json:"token"`
				ExpiresAt string `json:"expires_at"`
			} `json:"data"`
		}
		json.NewDecoder(resp.Body).Decode(&response)

		if response.Data.Token == "" || response.Data.UserID == 0 || response.Data.Username == "" || response.Data.Email == "" {
			t.Errorf("handler returned empty token")
		}

		if _, err := time.Parse(time.RFC3339, response.Data.ExpiresAt); err != nil {
			t.Errorf("handler returned invalid expires_at: %v", response.Data.ExpiresAt)
		}
	})
}
