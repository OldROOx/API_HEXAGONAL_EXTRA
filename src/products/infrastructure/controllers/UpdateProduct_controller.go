package controllers

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateProductController struct {
	updateProductUseCase *application.UpdateProductUseCase
}

func NewUpdateProductController(updateProductUseCase *application.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{updateProductUseCase: updateProductUseCase}
}

func (c *UpdateProductController) Handle(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := c.updateProductUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
