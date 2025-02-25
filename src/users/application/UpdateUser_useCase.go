package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type UpdateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUpdateUserUseCase(userRepo repositories.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: userRepo}
}

func (uc *UpdateUserUseCase) Execute(user *entities.User) error {
	return uc.userRepo.UpdateUser(user)
}
