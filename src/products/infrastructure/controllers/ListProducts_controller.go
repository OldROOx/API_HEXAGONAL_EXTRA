package controllers

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListProductsController struct {
	listProductsUseCase        *application.ListProductsUseCase
	getProductsByUserIDUseCase *application.GetProductsByUserIDUseCase
}

func NewListProductsController(
	listProductsUseCase *application.ListProductsUseCase,
	getProductsByUserIDUseCase *application.GetProductsByUserIDUseCase,
) *ListProductsController {
	return &ListProductsController{
		listProductsUseCase:        listProductsUseCase,
		getProductsByUserIDUseCase: getProductsByUserIDUseCase,
	}
}

func (c *ListProductsController) Handle(ctx *gin.Context) {
	// Check if we're filtering by user ID
	userID := ctx.Query("user_id")
	if userID != "" {
		id, err := strconv.ParseUint(userID, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}
		products, err := c.getProductsByUserIDUseCase.Execute(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, products)
		return
	}

	products, err := c.listProductsUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
