package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type HoSoRepository interface {
	Create(hoSo *entity.HoSo) error
	Update(hoSo *entity.HoSo) error
	Delete(id uint) error
	GetByID(id uint) (*entity.HoSo, error)
	GetAll() ([]*entity.HoSo, error)
}
type hoSoRepository struct {
	db *gorm.DB
}

func NewHoSoRepository(db *gorm.DB) HoSoRepository {
	return &hoSoRepository{db: db}
}
func (r *hoSoRepository) Create(hoSo *entity.HoSo) error {
	return r.db.Create(hoSo).Error
}
func (r *hoSoRepository) Update(hoSo *entity.HoSo) error {
	return r.db.Save(hoSo).Error
}
func (r *hoSoRepository) Delete(id uint) error {
	return r.db.Delete(&entity.HoSo{}, id).Error
}
func (r *hoSoRepository) GetByID(id uint) (*entity.HoSo, error) {
	var hoSo entity.HoSo
	if err := r.db.First(&hoSo, id).Error; err != nil {
		return nil, err
	}
	return &hoSo, nil
}
func (r *hoSoRepository) GetAll() ([]*entity.HoSo, error) {
	var hoSos []*entity.HoSo
	if err := r.db.Find(&hoSos).Error; err != nil {
		return nil, err
	}
	return hoSos, nil
}
