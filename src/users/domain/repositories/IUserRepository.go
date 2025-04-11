package repositories

import "API_HEXAGONAL_RECU/src/users/domain/entities"

type UserRepository interface {
	GetUsers() ([]entities.User, error)
	GetUsersWithProducts() ([]entities.UserWithProducts, error)
	GetUserByID(id uint) (*entities.User, error)
	GetUserWithProductsByID(id uint) (*entities.UserWithProducts, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id uint) error
}
