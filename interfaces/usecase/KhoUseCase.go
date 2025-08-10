package usecase

import (
	"context"
	"fmt"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
	"strconv"
)

type KhoUseCase interface {
	CreateKho(ctx context.Context, req *input.CreateKhoRequest) error
	GetAll(ctx context.Context) ([]entity.Kho, error)
	UpdateKho(ctx context.Context, id uint, req *input.UpdateKhoRequest) error
	DeleteKho(ctx context.Context, id uint) error
}

type khoUseCase struct {
	khoRepo repository.KhoRepository
}

func NewKhoUseCase(khoRepo repository.KhoRepository) KhoUseCase {
	return &khoUseCase{khoRepo: khoRepo}
}
func (k *khoUseCase) CreateKho(ctx context.Context, req *input.CreateKhoRequest) error {
	id, err := strconv.ParseInt(req.CoQuanID, 10, 64)
	if err != nil {
		return fmt.Errorf("CoQuanID không hợp lệ: %w", err)
	}
	kho := &entity.Kho{
		Ten:      req.Ten,
		CoQuanID: int(id),
	}
	return k.khoRepo.Create(kho)
}
func (k *khoUseCase) GetAll(ctx context.Context) ([]entity.Kho, error) {
	khos, err := k.khoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return khos, nil
}

func (k *khoUseCase) UpdateKho(ctx context.Context, id uint, req *input.UpdateKhoRequest) error {
	kho, err := k.khoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if kho.ID == 0 {
		return nil
	}
	kho.Ten = req.Ten
	return k.khoRepo.Update(&kho)
}
func (k *khoUseCase) DeleteKho(ctx context.Context, id uint) error {
	kho, err := k.khoRepo.GetByID(id)
	if err != nil {
		return err
	}
	if kho.ID == 0 {
		return nil // No need to delete if Kho does not exist
	}
	return k.khoRepo.Delete(id)
}
