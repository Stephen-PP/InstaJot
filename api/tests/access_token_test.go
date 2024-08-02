package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestRefreshTokenHandler(t *testing.T) {
	setupTests()

	t.Run("MethodNotAllowed", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:9183/auth/refresh-token", nil)
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
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/refresh-token", bytes.NewBuffer(body))
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

	t.Run("InvalidToken", func(t *testing.T) {
		body := map[string]string{"token": "invalid"}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/refresh-token", bytes.NewBuffer(bodyBytes))
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

	t.Run("ExpiredToken", func(t *testing.T) {
		expiredToken := "expiredtoken"
		db.Exec("INSERT INTO access_tokens (user_id, token, expires_at) VALUES (?, ?, ?)", 1, expiredToken, time.Now().Add(-1*time.Hour))

		body := map[string]string{"token": expiredToken}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/refresh-token", bytes.NewBuffer(bodyBytes))
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

	t.Run("SuccessfulTokenRefresh", func(t *testing.T) {
		validToken := "validtoken"
		db.Exec("INSERT INTO access_tokens (user_id, token, expires_at) VALUES (?, ?, ?)", 1, validToken, time.Now().Add(1*time.Hour))

		body := map[string]string{"token": validToken}
		bodyBytes, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:9183/auth/refresh-token", bytes.NewBuffer(bodyBytes))
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
				Token     string `json:"token"`
				ExpiresAt string `json:"expires_at"`
			}
		}
		json.NewDecoder(resp.Body).Decode(&response)

		if response.Data.Token == validToken {
			t.Errorf("handler returned the same token: got %v want a new token", response.Data.Token)
		}

		if response.Data.Token == "" {
			t.Errorf("handler returned an empty token: got %v want a new token", response.Data.Token)
		}

		if _, err := time.Parse(time.RFC3339, response.Data.ExpiresAt); err != nil {
			t.Errorf("handler returned invalid expires_at: %v", response.Data.ExpiresAt)
		}
	})
}
