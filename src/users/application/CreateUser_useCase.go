package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type CreateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewCreateUserUseCase(userRepo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) Execute(user *entities.User) error {
	return uc.userRepo.CreateUser(user)
}
