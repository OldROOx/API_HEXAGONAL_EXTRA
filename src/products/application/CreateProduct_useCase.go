package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type CreateProductUseCase struct {
	productRepo repositories.ProductRepository
}

func NewCreateProductUseCase(productRepo repositories.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{productRepo: productRepo}
}

func (uc *CreateProductUseCase) Execute(product *entities.Product) error {
	return uc.productRepo.CreateProduct(product)
}
