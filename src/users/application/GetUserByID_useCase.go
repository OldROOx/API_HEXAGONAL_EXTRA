package application

import (
	"API_HEXAGONAL_RECU/src/users/domain/entities"
	"API_HEXAGONAL_RECU/src/users/domain/repositories"
)

type GetUserByIDUseCase struct {
	userRepo repositories.UserRepository
}

func NewGetUserByIDUseCase(userRepo repositories.UserRepository) *GetUserByIDUseCase {
	return &GetUserByIDUseCase{userRepo: userRepo}
}

func (uc *GetUserByIDUseCase) Execute(id uint) (*entities.User, error) {
	return uc.userRepo.GetUserByID(id)
}

func (uc *GetUserByIDUseCase) ExecuteWithProducts(id uint) (*entities.UserWithProducts, error) {
	return uc.userRepo.GetUserWithProductsByID(id)
}
