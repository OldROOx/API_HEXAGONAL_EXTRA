package main

import (
	"API_HEXAGONAL_RECU/src/core/database"
	"API_HEXAGONAL_RECU/src/core/router"
	"API_HEXAGONAL_RECU/src/core/server"
	"API_HEXAGONAL_RECU/src/products/application"
	productController "API_HEXAGONAL_RECU/src/products/infrastructure/controllers"
	productRepo "API_HEXAGONAL_RECU/src/products/infrastructure/repositories/mysql"
	userApplication "API_HEXAGONAL_RECU/src/users/application"
	userController "API_HEXAGONAL_RECU/src/users/infrastructure/controllers"
	userRepo "API_HEXAGONAL_RECU/src/users/infrastructure/repositories/mysql"
	"log"
)

func main() {
	// Inicializar la conexión a la base de datos
	db := database.NewMySQLConnection()

	// Inicializar repositorios
	userRepository := userRepo.NewMySQLUserRepository(db)
	productRepository := productRepo.NewMySQLProductRepository(db)

	// Inicializar servicios
	userService := userApplication.NewUserService(userRepository)
	productService := application.NewProductService(productRepository)

	// Inicializar controladores
	userHandler := userController.NewUserController(userService)
	productHandler := productController.NewProductController(productService)

	// Configurar el router
	r := router.NewRouter(userHandler, productHandler)
	r.SetupRoutes()

	// Inicializar y arrancar el servidor
	srv := server.NewServer("8080", r.GetEngine())
	if err := srv.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
