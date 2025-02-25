package controllers

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetProductByIDController struct {
	getProductByIDUseCase *application.GetProductByIDUseCase
}

func NewGetProductByIDController(getProductByIDUseCase *application.GetProductByIDUseCase) *GetProductByIDController {
	return &GetProductByIDController{getProductByIDUseCase: getProductByIDUseCase}
}

func (c *GetProductByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := c.getProductByIDUseCase.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
