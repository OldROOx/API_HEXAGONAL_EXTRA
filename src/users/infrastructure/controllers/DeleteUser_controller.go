package controllers

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteUserController struct {
	deleteUserUseCase *application.DeleteUserUseCase
}

func NewDeleteUserController(deleteUserUseCase *application.DeleteUserUseCase) *DeleteUserController {
	return &DeleteUserController{deleteUserUseCase: deleteUserUseCase}
}

func (c *DeleteUserController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.deleteUserUseCase.Execute(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
