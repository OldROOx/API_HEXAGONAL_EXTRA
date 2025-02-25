package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type DeleteUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewDeleteUserUseCase(userRepo repositories.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: userRepo}
}

func (uc *DeleteUserUseCase) Execute(id uint) error {
	return uc.userRepo.DeleteUser(id)
}
