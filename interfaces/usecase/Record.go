package usecase

import (
	"context"
	"fmt"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
	"strconv"
)

type RecordUseCase interface {
	CreateRecord(ctx context.Context, req *input.CreateRecordRequest) error
	GetRecordByID(ctx context.Context, id uint) (*entity.Record, error)

	GetAllRecord(ctx context.Context) ([]*entity.Record, error)
	UpdateRecord(ctx context.Context, id uint, req *input.UpdateRecordRequest) error
	DeleteRecord(ctx context.Context, id uint) error
}
type RecordUseCaseImpl struct {
	recordRepo repository.RecordRepository
}

func NewRecordUseCase(recordRepo repository.RecordRepository) RecordUseCase {
	return &RecordUseCaseImpl{recordRepo: recordRepo}
}
func (h *RecordUseCaseImpl) CreateRecord(ctx context.Context, req *input.CreateRecordRequest) error {
	id, err := strconv.ParseInt(req.WarehouseID, 10, 64)
	if err != nil {
		return fmt.Errorf("Record không hợp lệ: %w", err)
	}
	record := &entity.Record{

		Name:            req.WarehouseID,
		WarehouseID:     int(id),
		ArchiveAgencyID: req.ArchiveAgencyID,
	}
	return h.recordRepo.Create(record)
}
func (h *RecordUseCaseImpl) GetRecordByID(ctx context.Context, id uint) (*entity.Record, error) {
	return h.recordRepo.GetByID(id)
}
func (h *RecordUseCaseImpl) GetAllRecord(ctx context.Context) ([]*entity.Record, error) {
	record, err := h.recordRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *RecordUseCaseImpl) UpdateRecord(ctx context.Context, id uint, req *input.UpdateRecordRequest) error {
	record, err := h.recordRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record.ID == 0 {
		return nil // No need to update if HoSo does not exist
	}
	record.Name = req.Name
	record.ArchiveAgencyID = req.ArchiveAgencyID
	return h.recordRepo.Update(record)
}
func (h *RecordUseCaseImpl) DeleteRecord(ctx context.Context, id uint) error {
	record, err := h.recordRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record.ID == 0 {
		return nil // No need to delete if HoSo does not exist
	}
	return h.recordRepo.Delete(id)
}
