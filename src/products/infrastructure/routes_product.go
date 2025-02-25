package infrastructure

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"API_HEXAGONAL_RECU/src/products/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(api *gin.RouterGroup, db *gorm.DB) {
	productRepo := NewMySQLProductRepository(db)

	// Use cases
	createProductUseCase := application.NewCreateProductUseCase(productRepo)
	deleteProductUseCase := application.NewDeleteProductUseCase(productRepo)
	getProductByIDUseCase := application.NewGetProductByIDUseCase(productRepo)
	listProductsUseCase := application.NewListProductsUseCase(productRepo)
	getProductsByUserIDUseCase := application.NewGetProductsByUserIDUseCase(productRepo)
	updateProductUseCase := application.NewUpdateProductUseCase(productRepo)

	// Controllers
	createProductController := controllers.NewCreateProductController(createProductUseCase)
	deleteProductController := controllers.NewDeleteProductController(deleteProductUseCase)
	getProductByIDController := controllers.NewGetProductByIDController(getProductByIDUseCase)
	listProductsController := controllers.NewListProductsController(listProductsUseCase, getProductsByUserIDUseCase)
	updateProductController := controllers.NewUpdateProductController(updateProductUseCase)

	// Routes
	products := api.Group("/products")
	{
		products.GET("", listProductsController.Handle)
		products.POST("", createProductController.Handle)
		products.GET("/:id", getProductByIDController.Handle)
		products.PUT("/:id", updateProductController.Handle)
		products.DELETE("/:id", deleteProductController.Handle)
	}
}
