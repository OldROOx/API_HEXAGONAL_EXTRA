package controllers

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteProductController struct {
	deleteProductUseCase *application.DeleteProductUseCase
}

func NewDeleteProductController(deleteProductUseCase *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{deleteProductUseCase: deleteProductUseCase}
}

func (c *DeleteProductController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := c.deleteProductUseCase.Execute(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
