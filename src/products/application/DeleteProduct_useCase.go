package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type DeleteProductUseCase struct {
	productRepo repositories.ProductRepository
}

func NewDeleteProductUseCase(productRepo repositories.ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{productRepo: productRepo}
}

func (uc *DeleteProductUseCase) Execute(id uint) error {
	return uc.productRepo.DeleteProduct(id)
}
