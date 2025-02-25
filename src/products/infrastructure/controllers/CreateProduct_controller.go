package controllers

import (
	"API_HEXAGONAL_RECU/src/products/application"
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductController struct {
	createProductUseCase *application.CreateProductUseCase
}

func NewCreateProductController(createProductUseCase *application.CreateProductUseCase) *CreateProductController {
	return &CreateProductController{createProductUseCase: createProductUseCase}
}

func (c *CreateProductController) Handle(ctx *gin.Context) {
	var product entities.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.createProductUseCase.Execute(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
