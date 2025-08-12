package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	GetByUserName(user string) (*entity.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByUserName(user string) (*entity.User, error) {
	var u entity.User
	if err := r.db.Where("username = ?", user).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
