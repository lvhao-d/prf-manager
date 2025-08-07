package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
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
