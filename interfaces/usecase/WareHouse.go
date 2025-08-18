package usecase

import (
	"context"
	"fmt"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
	"strconv"
)

type WareHouseUseCase interface {
	CreateWareHouse(ctx context.Context, req *input.CreateWareHouseRequest) error
	GetAll(ctx context.Context) ([]entity.Warehouse, error)
	UpdateWareHouse(ctx context.Context, id uint, req *input.UpdateWareHouseRequest) error
	DeleteWareHouse(ctx context.Context, id uint) error
}

type wareHouseUseCase struct {
	wareHouseRepo repository.WareHouseRepository
}

func NewWareHouseUseCase(wareHouseRepo repository.WareHouseRepository) WareHouseUseCase {
	return &wareHouseUseCase{wareHouseRepo: wareHouseRepo}
}
func (k *wareHouseUseCase) CreateWareHouse(ctx context.Context, req *input.CreateWareHouseRequest) error {
	id, err := strconv.ParseInt(req.AgencyID, 10, 64)
	if err != nil {
		return fmt.Errorf("AgencyID không hợp lệ: %w", err)
	}
	wareHouse := &entity.Warehouse{
		Name:     req.Name,
		AgencyID: int(id),
	}
	return k.wareHouseRepo.Create(wareHouse)
}
func (k *wareHouseUseCase) GetAll(ctx context.Context) ([]entity.Warehouse, error) {
	wareHouse, err := k.wareHouseRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return wareHouse, nil
}

func (k *wareHouseUseCase) UpdateWareHouse(ctx context.Context, id uint, req *input.UpdateWareHouseRequest) error {
	wareHouse, err := k.wareHouseRepo.GetByID(id)
	if err != nil {
		return err
	}
	if wareHouse.ID == 0 {
		return nil
	}
	wareHouse.Name = req.Name
	return k.wareHouseRepo.Update(&wareHouse)
}
func (k *wareHouseUseCase) DeleteWareHouse(ctx context.Context, id uint) error {
	wareHouse, err := k.wareHouseRepo.GetByID(id)
	if err != nil {
		return err
	}
	if wareHouse.ID == 0 {
		return nil
	}
	return k.wareHouseRepo.Delete(id)
}
