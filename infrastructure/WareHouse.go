package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type WareHouseRepository interface {
	Create(wareHouse *entity.Warehouse) error
	GetAll() ([]entity.Warehouse, error)
	GetByID(id uint) (entity.Warehouse, error)
	Update(wareHouse *entity.Warehouse) error
	Delete(id uint) error
}

type wareHouseRepository struct {
	db *gorm.DB
}

func NewWareHouseRepository(db *gorm.DB) WareHouseRepository {
	return &wareHouseRepository{db: db}
}
func (r *wareHouseRepository) Create(wareHouse *entity.Warehouse) error {
	return r.db.
		Create(wareHouse).Error
}
func (r *wareHouseRepository) GetAll() ([]entity.Warehouse, error) {
	var wareHouse []entity.Warehouse
	if err := r.db.
		Preload("Records").
		Find(&wareHouse).Error; err != nil {
		return nil, err
	}
	return wareHouse, nil
}
func (r *wareHouseRepository) GetByID(id uint) (entity.Warehouse, error) {
	var wareHouse entity.Warehouse
	if err := r.db.
		First(&wareHouse, id).Error; err != nil {
		return entity.Warehouse{}, err
	}
	return wareHouse, nil
}

func (r *wareHouseRepository) Update(wareHouse *entity.Warehouse) error {
	return r.db.
		Save(wareHouse).Error
}
func (r *wareHouseRepository) Delete(id uint) error {
	return r.db.
		Delete(&entity.Warehouse{}, id).Error
}
