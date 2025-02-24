package application

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/products/domain/repositories"
)

type ProductService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetProducts() ([]entities.Product, error) {
	return s.productRepo.GetProducts()
}

func (s *ProductService) GetProductByID(id uint) (*entities.Product, error) {
	return s.productRepo.GetProductByID(id)
}

func (s *ProductService) GetProductsByUserID(userID uint) ([]entities.Product, error) {
	return s.productRepo.GetProductsByUserID(userID)
}

func (s *ProductService) CreateProduct(product *entities.Product) error {
	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *entities.Product) error {
	return s.productRepo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.productRepo.DeleteProduct(id)
}
