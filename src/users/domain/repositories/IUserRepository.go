package repositories

import "API_HEXAGONAL_RECU/src/users/domain/entities"

type UserRepository interface {
	GetUsers() ([]entities.User, error)
	GetUserByID(id uint) (*entities.User, error)
	CreateUser(user *entities.User) error
	UpdateUser(user *entities.User) error
	DeleteUser(id uint) error
}
