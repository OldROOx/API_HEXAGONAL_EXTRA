package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type ListProductsUseCase struct {
	productRepo repositories.ProductRepository
}

func NewListProductsUseCase(productRepo repositories.ProductRepository) *ListProductsUseCase {
	return &ListProductsUseCase{productRepo: productRepo}
}

func (uc *ListProductsUseCase) Execute() ([]entities.Product, error) {
	return uc.productRepo.GetProducts()
}
