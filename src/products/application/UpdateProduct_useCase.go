package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type UpdateProductUseCase struct {
	productRepo repositories.ProductRepository
}

func NewUpdateProductUseCase(productRepo repositories.ProductRepository) *UpdateProductUseCase {
	return &UpdateProductUseCase{productRepo: productRepo}
}

func (uc *UpdateProductUseCase) Execute(product *entities.Product) error {
	return uc.productRepo.UpdateProduct(product)
}
