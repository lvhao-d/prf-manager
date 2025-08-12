package repository

import (
	"prf-manager/entity"

	"gorm.io/gorm"
)

type RecordRepository interface {
	Create(record *entity.Record) error
	Update(record *entity.Record) error
	Delete(id uint) error
	GetByID(id uint) (*entity.Record, error)
	GetAll() ([]*entity.Record, error)
}
type recordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{db: db}
}
func (r *recordRepository) Create(record *entity.Record) error {
	return r.db.
		Create(record).Error
}
func (r *recordRepository) Update(record *entity.Record) error {
	return r.db.
		Save(record).Error
}
func (r *recordRepository) Delete(id uint) error {
	return r.db.
		Delete(&entity.Record{}, id).Error
}
func (r *recordRepository) GetByID(id uint) (*entity.Record, error) {
	var record entity.Record
	if err := r.db.
		First(&record, id).Error; err != nil {
		return nil, err
	}
	return &record, nil
}
func (r *recordRepository) GetAll() ([]*entity.Record, error) {
	var record []*entity.Record
	if err := r.db.
		Find(&record).Error; err != nil {
		return nil, err
	}
	return record, nil
}
