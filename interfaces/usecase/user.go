package usecase

import (
	"context"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
	"prf-manager/interfaces/output"
	"prf-manager/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Create(ctx context.Context, p *input.CreateUserRequest) error
	Login(ctx context.Context, p *input.UserLoginRequest) (*output.LoginResponse, error)
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
		Password: p.Password,
	}

	return u.userRepo.Create(user)
}

func (u *userUseCase) Login(ctx context.Context, p *input.UserLoginRequest) (*output.LoginResponse, error) {
	user, err := u.userRepo.GetByUserName(p.Username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &output.LoginResponse{
		ID:          user.ID,
		Username:    user.Username,
		AccessToken: accessToken,
	}, nil

}
