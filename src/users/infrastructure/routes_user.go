package infrastructure

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"API_HEXAGONAL_RECU/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(api *gin.RouterGroup, db *gorm.DB) {
	userRepo := NewMySQLUserRepository(db)

	// Use cases
	createUserUseCase := application.NewCreateUserUseCase(userRepo)
	deleteUserUseCase := application.NewDeleteUserUseCase(userRepo)
	getUserByIDUseCase := application.NewGetUserByIDUseCase(userRepo)
	listUsersUseCase := application.NewListUsersUseCase(userRepo)
	updateUserUseCase := application.NewUpdateUserUseCase(userRepo)

	// Controllers
	createUserController := controllers.NewCreateUserController(createUserUseCase)
	deleteUserController := controllers.NewDeleteUserController(deleteUserUseCase)
	getUserByIDController := controllers.NewGetUserByIDController(getUserByIDUseCase)
	listUsersController := controllers.NewListUsersController(listUsersUseCase)
	updateUserController := controllers.NewUpdateUserController(updateUserUseCase)

	// Routes
	users := api.Group("/users")
	{
		users.GET("", listUsersController.Handle)
		users.POST("", createUserController.Handle)
		users.GET("/:id", getUserByIDController.Handle)
		users.PUT("/:id", updateUserController.Handle)
		users.DELETE("/:id", deleteUserController.Handle)
	}
}
