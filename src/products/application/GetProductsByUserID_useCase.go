package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type GetProductsByUserIDUseCase struct {
	productRepo repositories.ProductRepository
}

func NewGetProductsByUserIDUseCase(productRepo repositories.ProductRepository) *GetProductsByUserIDUseCase {
	return &GetProductsByUserIDUseCase{productRepo: productRepo}
}

func (uc *GetProductsByUserIDUseCase) Execute(userID uint) ([]entities.Product, error) {
	return uc.productRepo.GetProductsByUserID(userID)
}
