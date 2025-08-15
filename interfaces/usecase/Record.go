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
	TransferToArchive(ctx context.Context, id uint, req *input.UpdateRecordRequest) error
	GetAllRecord(ctx context.Context) ([]*entity.Record, error)
	UpdateRecord(ctx context.Context, id uint, req *input.UpdateRecordRequest) error
	DeleteRecord(ctx context.Context, id uint) error
	SearchRecord(ctx context.Context, req *input.Search) ([]*entity.Record, error)
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
		return fmt.Errorf("%w", err)
	}
	ArchiveAgencyID, _ := strconv.ParseInt(req.ArchiveAgencyID, 10, 64)
	record := &entity.Record{

		Name:            req.Name,
		WarehouseID:     int(id),
		ArchiveAgencyID: int(ArchiveAgencyID),
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
		return nil
	}
	ArchiveAgencyID, _ := strconv.ParseInt(req.ArchiveAgencyID, 10, 64)
	WarehouseID, _ := strconv.ParseInt(req.WarehouseID, 10, 64)
	record.WarehouseID = int(WarehouseID)
	record.Name = req.Name
	record.ArchiveAgencyID = int(ArchiveAgencyID)

	return h.recordRepo.Update(record)
}
func (h *RecordUseCaseImpl) DeleteRecord(ctx context.Context, id uint) error {
	record, err := h.recordRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record.ID == 0 {
		return nil
	}
	return h.recordRepo.Delete(id)
}

func (h *RecordUseCaseImpl) TransferToArchive(ctx context.Context, id uint, req *input.UpdateRecordRequest) error {
	record, err := h.recordRepo.GetByID(id)
	if err != nil {
		return err
	}
	if record.ID == 0 {
		return nil
	}
	WarehouseID, _ := strconv.ParseInt(req.WarehouseID, 10, 64)
	record.WarehouseID = int(WarehouseID)
	record.Name = req.Name
	record.ArchiveAgencyID = 99

	return h.recordRepo.Update(record)
}

func (h *RecordUseCaseImpl) SearchRecord(ctx context.Context, req *input.Search) ([]*entity.Record, error) {
	warehouseID, _ := strconv.Atoi(req.WarehouseID)
	archiveAgencyID, _ := strconv.Atoi(req.ArchiveAgencyID)

	return h.recordRepo.SearchRecord(warehouseID, archiveAgencyID)
}
