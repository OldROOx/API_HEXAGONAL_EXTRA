package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUsers() ([]entities.User, error) {
	return s.userRepo.GetUsers()
}

func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) UpdateUser(user *entities.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
