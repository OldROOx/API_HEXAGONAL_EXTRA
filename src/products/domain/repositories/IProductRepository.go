package repositories

import "API_HEXAGONAL_RECU/src/products/domain/entities"

type ProductRepository interface {
	GetProducts() ([]entities.Product, error)
	GetProductByID(id uint) (*entities.Product, error)
	GetProductsByUserID(userID uint) ([]entities.Product, error)
	CreateProduct(product *entities.Product) error
	UpdateProduct(product *entities.Product) error
	DeleteProduct(id uint) error
}
