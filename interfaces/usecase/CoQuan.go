package usecase

import (
	"context"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
)

type CoQuanUseCase interface {
	CreateCoQuan(ctx context.Context, p *input.CreateCoQuanRequest) error
	UpdateCoQuan(ctx context.Context, id uint, p *input.UpdateCoQuanRequest) error
	DeleteCoQuan(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]entity.CoQuan, error)
}

type coQuanUseCase struct {
	coQuanRepo repository.CoQuanRepository
}

func NewCoQuanUseCase(coQuanRepo repository.CoQuanRepository) CoQuanUseCase {
	return &coQuanUseCase{coQuanRepo: coQuanRepo}
}
func (c *coQuanUseCase) CreateCoQuan(ctx context.Context, p *input.CreateCoQuanRequest) error {
	coQuan := &entity.CoQuan{
		Ten: p.Name,
	}
	return c.coQuanRepo.Create(coQuan)

}

func (c *coQuanUseCase) GetAll(ctx context.Context) ([]entity.CoQuan, error) {
	return c.coQuanRepo.GetAll()
}

func (c *coQuanUseCase) UpdateCoQuan(ctx context.Context, id uint, p *input.UpdateCoQuanRequest) error {
	coQuan, err := c.coQuanRepo.GetByID(id)
	if err != nil {
		return err
	}
	coQuan.Ten = p.Name
	return c.coQuanRepo.Update(&coQuan)
}
func (c *coQuanUseCase) DeleteCoQuan(ctx context.Context, id uint) error {
	coQuan, err := c.coQuanRepo.GetByID(id)
	if err != nil {
		return err
	}
	if coQuan.ID == 0 {
		return nil // No need to delete if CoQuan does not exist
	}
	return c.coQuanRepo.Delete(id)
}
