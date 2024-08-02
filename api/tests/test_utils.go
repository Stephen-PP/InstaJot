package tests

import (
	"database/sql"

	"github.com/stephen-pp/instajot/internal/api"
	"github.com/stephen-pp/instajot/internal/database"
)

var db *sql.DB
var server *api.APIServer

func setupTests() *sql.DB {
	if db == nil {
		err := database.InitDB(":memory:")
		if err != nil {
			panic(err)
		}

		db = database.GetDB()
	}

	if server == nil {
		server = &api.APIServer{DB: db}
		go func() {
			server.Start("", "9183")
		}()
	}

	return db
}
