package usecase

import (
	"context"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
)

type UserUseCase interface {
	Create(ctx context.Context, p *input.CreateUserRequest) error
}
type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) Create(ctx context.Context, p *input.CreateUserRequest) error {
	user := &entity.User{
		Username: p.Username,
		Email:    p.Email,
		Password: p.Password,
	}

	return u.userRepo.Create(user)
}
