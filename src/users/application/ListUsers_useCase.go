package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type ListUsersUseCase struct {
	userRepo repositories.UserRepository
}

func NewListUsersUseCase(userRepo repositories.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepo: userRepo}
}

func (uc *ListUsersUseCase) Execute() ([]entities.User, error) {
	return uc.userRepo.GetUsers()
}
