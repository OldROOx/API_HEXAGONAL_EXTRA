package controllers

import (
	"API_HEXAGONAL_RECU/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetUserByIDController struct {
	getUserByIDUseCase *application.GetUserByIDUseCase
}

func NewGetUserByIDController(getUserByIDUseCase *application.GetUserByIDUseCase) *GetUserByIDController {
	return &GetUserByIDController{getUserByIDUseCase: getUserByIDUseCase}
}

func (c *GetUserByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := c.getUserByIDUseCase.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
