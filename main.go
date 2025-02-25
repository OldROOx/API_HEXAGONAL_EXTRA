package main

import (
	"API_HEXAGONAL_RECU/src/core/database"
	"API_HEXAGONAL_RECU/src/core/router"
	"API_HEXAGONAL_RECU/src/core/server"
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	userEntities "API_HEXAGONAL_RECU/src/users/domain/entities"
	"log"
)

func main() {
	// Database connection
	db := database.NewMySQLConnection()

	// Auto-migrate entities
	err := db.AutoMigrate(&userEntities.User{}, &entities.Product{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Router setup
	r := router.NewRouter(db)
	r.SetupRoutes()

	// Create server
	srv := server.NewServer("8080", r.GetEngine())

	// Start server
	log.Fatal(srv.Start())
}
