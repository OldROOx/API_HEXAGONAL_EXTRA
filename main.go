package main

import (
	"API_HEXAGONAL_RECU/src/core/database"
	"API_HEXAGONAL_RECU/src/core/router"
	"API_HEXAGONAL_RECU/src/core/server"
	"log"
)

func main() {
	// Database connection
	db := database.NewMySQLConnection()

	// Skip auto-migration to preserve existing data
	// We'll update the schema manually if needed

	// Router setup
	r := router.NewRouter(db)
	r.SetupRoutes()

	// Create server
	srv := server.NewServer("8080", r.GetEngine())

	// Start server
	log.Fatal(srv.Start())
}
