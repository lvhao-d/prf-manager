package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type AgencyRepository interface {
	Create(agency *entity.Agency) error
	GetByID(id uint) (entity.Agency, error)
	Update(agency *entity.Agency) error
	Delete(id uint) error

	GetAll() ([]entity.Agency, error)
}

type agencyRepository struct {
	db *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) AgencyRepository {
	return &agencyRepository{db: db}
}
func (r *agencyRepository) Create(agency *entity.Agency) error {
	return r.db.
		Create(agency).Error
}
func (r *agencyRepository) GetByID(id uint) (entity.Agency, error) {
	var agency entity.Agency
	if err := r.db.First(&agency, id).Error; err != nil {
		return entity.Agency{}, err
	}
	return agency, nil
}
func (r *agencyRepository) Update(agency *entity.Agency) error {
	return r.db.
		Save(agency).Error
}
func (r *agencyRepository) Delete(id uint) error {
	return r.db.
		Delete(&entity.Agency{}, id).Error
}
func (r *agencyRepository) GetAll() ([]entity.Agency, error) {
	var agency []entity.Agency
	if err := r.db.
		Preload("Warehouses").
		Preload("Warehouses.Records").
		Find(&agency).Error; err != nil {
		return nil, err
	}
	return agency, nil
}
