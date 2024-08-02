package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	DB     *sql.DB
	Server *http.Server
}

func (api *APIServer) Start(addr string, port string) {
	mux := http.NewServeMux()
	// Register routes
	mux.HandleFunc("/auth/register", api.registerHandler)
	mux.HandleFunc("/auth/login", api.loginHandler)
	mux.HandleFunc("/auth/refresh-token", api.refreshTokenHandler)

	api.Server = &http.Server{
		Addr:    addr + ":" + port,
		Handler: mux,
	}

	err := api.Server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (api *APIServer) Shutdown() error {
	return api.Server.Shutdown(context.Background())
}
