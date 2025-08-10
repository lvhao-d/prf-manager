package usecase

import (
	"context"
	"fmt"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
	"strconv"
)

type HoSoUseCase interface {
	CreateHoSo(ctx context.Context, req *input.CreateHoSoRequest) error
	GetHoSoByID(ctx context.Context, id uint) (*entity.HoSo, error)

	GetAllHoSo(ctx context.Context) ([]*entity.HoSo, error)
	UpdateHoSo(ctx context.Context, id uint, req *input.UpdateHoSoRequest) error
	DeleteHoSo(ctx context.Context, id uint) error
}
type HoSoUseCaseImpl struct {
	hoSoRepo repository.HoSoRepository
}

func NewHoSoUseCase(hoSoRepo repository.HoSoRepository) HoSoUseCase {
	return &HoSoUseCaseImpl{hoSoRepo: hoSoRepo}
}
func (h *HoSoUseCaseImpl) CreateHoSo(ctx context.Context, req *input.CreateHoSoRequest) error {
	id, err := strconv.ParseInt(req.KhoID, 10, 64)
	if err != nil {
		return fmt.Errorf("kho không hợp lệ: %w", err)
	}
	hoSo := &entity.HoSo{

		Ten:            req.Ten,
		KhoID:          int(id),
		CoQuanLuuTruID: req.CoQuanLuuTruID,
	}
	return h.hoSoRepo.Create(hoSo)
}
func (h *HoSoUseCaseImpl) GetHoSoByID(ctx context.Context, id uint) (*entity.HoSo, error) {
	return h.hoSoRepo.GetByID(id)
}
func (h *HoSoUseCaseImpl) GetAllHoSo(ctx context.Context) ([]*entity.HoSo, error) {
	hoSos, err := h.hoSoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return hoSos, nil
}

func (h *HoSoUseCaseImpl) UpdateHoSo(ctx context.Context, id uint, req *input.UpdateHoSoRequest) error {
	hoSo, err := h.hoSoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if hoSo.ID == 0 {
		return nil // No need to update if HoSo does not exist
	}
	hoSo.Ten = req.Ten
	hoSo.CoQuanLuuTruID = req.CoQuanLuuTruID
	return h.hoSoRepo.Update(hoSo)
}
func (h *HoSoUseCaseImpl) DeleteHoSo(ctx context.Context, id uint) error {
	hoSo, err := h.hoSoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if hoSo.ID == 0 {
		return nil // No need to delete if HoSo does not exist
	}
	return h.hoSoRepo.Delete(id)
}
