package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type CoQuanRepository interface {
	Create(coQuan *entity.CoQuan) error
	GetByID(id uint) (entity.CoQuan, error)
	Update(coQuan *entity.CoQuan) error
	Delete(id uint) error

	GetAll() ([]entity.CoQuan, error)
}

type coQuanRepository struct {
	db *gorm.DB
}

func NewCoQuanRepository(db *gorm.DB) CoQuanRepository {
	return &coQuanRepository{db: db}
}
func (r *coQuanRepository) Create(coQuan *entity.CoQuan) error {
	return r.db.Create(coQuan).Error
}
func (r *coQuanRepository) GetByID(id uint) (entity.CoQuan, error) {
	var coQuan entity.CoQuan
	if err := r.db.First(&coQuan, id).Error; err != nil {
		return entity.CoQuan{}, err
	}
	return coQuan, nil
}
func (r *coQuanRepository) Update(coQuan *entity.CoQuan) error {
	return r.db.Save(coQuan).Error
}
func (r *coQuanRepository) Delete(id uint) error {
	return r.db.Delete(&entity.CoQuan{}, id).Error
}
func (r *coQuanRepository) GetAll() ([]entity.CoQuan, error) {
	var coQuans []entity.CoQuan
	if err := r.db.Find(&coQuans).Error; err != nil {
		return nil, err
	}
	return coQuans, nil
}
