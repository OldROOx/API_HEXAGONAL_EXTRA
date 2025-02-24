package mysql

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) GetUsers() ([]entities.User, error) {
	var users []entities.User
	result := r.db.Preload("Products").Find(&users)
	return users, result.Error
}

func (r *MySQLUserRepository) GetUserByID(id uint) (*entities.User, error) {
	var user entities.User
	result := r.db.Preload("Products").First(&user, id)
	return &user, result.Error
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
