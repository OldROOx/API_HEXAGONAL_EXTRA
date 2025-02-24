package router

import (
	productController "API_HEXAGONAL_RECU/src/products/infrastructure/controllers"
	userController "API_HEXAGONAL_RECU/src/users/infrastructure/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Router struct {
	engine            *gin.Engine
	userController    *userController.UserController
	productController *productController.ProductController
}

func NewRouter(
	userController *userController.UserController,
	productController *productController.ProductController,
) *Router {
	engine := gin.Default()

	// Configuración CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Router{
		engine:            engine,
		userController:    userController,
		productController: productController,
	}
}

func (r *Router) SetupRoutes() {
	api := r.engine.Group("/api")
	{
		// Rutas de usuarios
		users := api.Group("/users")
		{
			users.GET("", r.userController.GetUsers)
			users.POST("", r.userController.CreateUser)
			users.GET("/:id", r.userController.GetUserByID)
			users.PUT("/:id", r.userController.UpdateUser)
			users.DELETE("/:id", r.userController.DeleteUser)
		}

		// Rutas de productos
		products := api.Group("/products")
		{
			products.GET("", r.productController.GetProducts)
			products.POST("", r.productController.CreateProduct)
			products.GET("/:id", r.productController.GetProductByID)
			products.PUT("/:id", r.productController.UpdateProduct)
			products.DELETE("/:id", r.productController.DeleteProduct)
		}
	}
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
