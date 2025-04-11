package infrastructure

import (
	productEntities "API_HEXAGONAL_RECU/src/products/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"gorm.io/gorm"
	"sync"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

// Singleton pattern implementation
var (
	userRepoInstance *MySQLUserRepository
	userRepoOnce     sync.Once
)

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	userRepoOnce.Do(func() {
		userRepoInstance = &MySQLUserRepository{db: db}
	})
	return userRepoInstance
}

func (r *MySQLUserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *MySQLUserRepository) GetUsersWithProducts() ([]entities.UserWithProducts, error) {
	// First get all users
	var users []entities.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	// Then get products and build the DTOs
	var usersWithProducts []entities.UserWithProducts

	for _, user := range users {
		// Get products for this user
		var products []productEntities.Product
		if err := r.db.Where("user_id = ?", user.ID).Find(&products).Error; err != nil {
			return nil, err
		}

		// Convert to product associations
		productAssociations := make([]entities.ProductAssociation, len(products))
		for i, product := range products {
			productAssociations[i] = entities.ProductAssociation{
				ID:     product.ID,
				Name:   product.Name,
				Price:  product.Price,
				UserID: product.UserID,
			}
		}

		// Create DTO
		userWithProducts := user.ToUserWithProducts(productAssociations)
		usersWithProducts = append(usersWithProducts, *userWithProducts)
	}

	return usersWithProducts, nil
}

func (r *MySQLUserRepository) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *MySQLUserRepository) GetUserWithProductsByID(id uint) (*entities.UserWithProducts, error) {
	// Get user
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Get products
	var products []productEntities.Product
	if err := r.db.Where("user_id = ?", user.ID).Find(&products).Error; err != nil {
		return nil, err
	}

	// Convert to product associations
	productAssociations := make([]entities.ProductAssociation, len(products))
	for i, product := range products {
		productAssociations[i] = entities.ProductAssociation{
			ID:     product.ID,
			Name:   product.Name,
			Price:  product.Price,
			UserID: product.UserID,
		}
	}

	// Create and return the DTO
	return user.ToUserWithProducts(productAssociations), nil
}

func (r *MySQLUserRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) UpdateUser(user *entities.User) error {
	return r.db.Save(user).Error
}

func (r *MySQLUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&entities.User{}, id).Error
}
