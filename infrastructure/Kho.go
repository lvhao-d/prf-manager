package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type KhoRepository interface {
	Create(kho *entity.Kho) error
	GetAll() ([]entity.Kho, error)
	GetByID(id uint) (entity.Kho, error)
	Update(kho *entity.Kho) error
	Delete(id uint) error
}

type khoRepository struct {
	db *gorm.DB
}

func NewKhoRepository(db *gorm.DB) KhoRepository {
	return &khoRepository{db: db}
}
func (r *khoRepository) Create(kho *entity.Kho) error {
	return r.db.Create(kho).Error
}
func (r *khoRepository) GetAll() ([]entity.Kho, error) {
	var khos []entity.Kho
	if err := r.db.Find(&khos).Error; err != nil {
		return nil, err
	}
	return khos, nil
}
func (r *khoRepository) GetByID(id uint) (entity.Kho, error) {
	var kho entity.Kho
	if err := r.db.First(&kho, id).Error; err != nil {
		return entity.Kho{}, err
	}
	return kho, nil
}

func (r *khoRepository) Update(kho *entity.Kho) error {
	return r.db.Save(kho).Error
}
func (r *khoRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Kho{}, id).Error
}
