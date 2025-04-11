package controllers

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListUsersController struct {
	listUsersUseCase *application.ListUsersUseCase
}

func NewListUsersController(listUsersUseCase *application.ListUsersUseCase) *ListUsersController {
	return &ListUsersController{listUsersUseCase: listUsersUseCase}
}

func (c *ListUsersController) Handle(ctx *gin.Context) {
	// Check if we should include products
	includeProducts := ctx.Query("include_products") == "true"

	if includeProducts {
		usersWithProducts, err := c.listUsersUseCase.ExecuteWithProducts()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, usersWithProducts)
	} else {
		users, err := c.listUsersUseCase.Execute()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, users)
	}
}
