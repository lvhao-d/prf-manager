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
	GetAll(page int) ([]*entity.Record, int64, error)
	SearchRecord(warehouseID, archiveAgencyID int) ([]*entity.Record, error)
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
func (r *recordRepository) GetAll(page int) ([]*entity.Record, int64, error) {
	if page < 1 {
		page = 1
	}
	limit := 5
	offset := (page - 1) * limit

	var record []*entity.Record

	var total int64

	if err := r.db.Model(&entity.Record{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.
		Offset(offset).
		Limit(limit).
		Find(&record).Error; err != nil {
		return nil, 0, err
	}
	return record, total, nil

	// // if err := r.db.
	// // 	Find(&record).Error; err != nil {
	// // 	return nil, err
	// // }
	// return record, nil
}

func (r *recordRepository) SearchRecord(warehouseID, archiveAgencyID int) ([]*entity.Record, error) {
	var records []*entity.Record

	db := r.db

	if archiveAgencyID != 0 {
		db = db.Where("archive_agency_id = ?", archiveAgencyID)
	} else {
		db.Find(&records)
	}

	if warehouseID != 0 {
		db = db.Where("warehouse_id = ?", warehouseID)
	}

	if err := db.Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
