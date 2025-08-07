package usecase

import (
	"context"
	"prf-manager/entity"
	"prf-manager/interfaces/input"
)

type KhoUseCase interface {
	CreateKho(ctx context.Context, req *input.CreateKhoRequest) error
	GetAll(ctx context.Context) ([]entity.Kho, error)
}
