package gorm

import (
	"prf-manager/internal/domain"

	"gorm.io/gorm"
)

type AgencyRepository struct {
	db *gorm.DB
}

func NewAgencyRepository(db *gorm.DB) *AgencyRepository {
	return &AgencyRepository{db: db}
}
func (r *AgencyRepository) GetAll() ([]domain.Agency, error) {
	var agencies []domain.Agency
	if err := r.db.Find(&agencies).Error; err != nil {
		return nil, err
	}
	return agencies, nil
}
func (r *AgencyRepository) Create(agency *domain.Agency) error {
	if err := r.db.Create(agency).Error; err != nil {
		return err
	}
	return nil
}
