package controllers

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserController struct {
	createUserUseCase *application.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase *application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{createUserUseCase: createUserUseCase}
}

func (c *CreateUserController) Handle(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.createUserUseCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
