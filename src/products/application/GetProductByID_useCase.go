package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type GetProductByIDUseCase struct {
	productRepo repositories.ProductRepository
}

func NewGetProductByIDUseCase(productRepo repositories.ProductRepository) *GetProductByIDUseCase {
	return &GetProductByIDUseCase{productRepo: productRepo}
}

func (uc *GetProductByIDUseCase) Execute(id uint) (*entities.Product, error) {
	return uc.productRepo.GetProductByID(id)
}
