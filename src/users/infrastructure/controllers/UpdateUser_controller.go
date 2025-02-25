package controllers

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateUserController struct {
	updateUserUseCase *application.UpdateUserUseCase
}

func NewUpdateUserController(updateUserUseCase *application.UpdateUserUseCase) *UpdateUserController {
	return &UpdateUserController{updateUserUseCase: updateUserUseCase}
}

func (c *UpdateUserController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)
	if err := c.updateUserUseCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
