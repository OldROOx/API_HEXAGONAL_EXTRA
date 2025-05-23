package infrastructure

import (
	"API_HEXAGONAL_RECU/src/products/domain/entities"
	"gorm.io/gorm"
	"sync"
)

type MySQLProductRepository struct {
	db *gorm.DB
}

// Singleton pattern implementation
var (
	productRepoInstance *MySQLProductRepository
	productRepoOnce     sync.Once
)

func NewMySQLProductRepository(db *gorm.DB) *MySQLProductRepository {
	productRepoOnce.Do(func() {
		productRepoInstance = &MySQLProductRepository{db: db}
	})
	return productRepoInstance
}

func (r *MySQLProductRepository) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.Find(&products)
	return products, result.Error
}

func (r *MySQLProductRepository) GetProductByID(id uint) (*entities.Product, error) {
	var product entities.Product
	result := r.db.First(&product, id)
	return &product, result.Error
}

func (r *MySQLProductRepository) GetProductsByUserID(userID uint) ([]entities.Product, error) {
	var products []entities.Product
	result := r.db.Where("user_id = ?", userID).Find(&products)
	return products, result.Error
}

func (r *MySQLProductRepository) CreateProduct(product *entities.Product) error {
	return r.db.Create(product).Error
}

func (r *MySQLProductRepository) UpdateProduct(product *entities.Product) error {
	return r.db.Save(product).Error
}

func (r *MySQLProductRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&entities.Product{}, id).Error
}
