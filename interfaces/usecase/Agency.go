package usecase

import (
	"context"
	"prf-manager/entity"
	repository "prf-manager/infrastructure"
	"prf-manager/interfaces/input"
)

type AgencyUseCase interface {
	CreateAgency(ctx context.Context, p *input.CreateAgencyRequest) error
	UpdateAgency(ctx context.Context, id uint, p *input.UpdateAgencyRequest) error
	DeleteAgency(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]entity.Agency, error)
}

type agencyUseCase struct {
	agencyRepo repository.AgencyRepository
}

func NewAgencyUseCase(agencyRepo repository.AgencyRepository) AgencyUseCase {
	return &agencyUseCase{agencyRepo: agencyRepo}
}
func (c *agencyUseCase) CreateAgency(ctx context.Context, p *input.CreateAgencyRequest) error {
	agency := &entity.Agency{
		Name: p.Name,
	}
	return c.agencyRepo.Create(agency)

}

func (c *agencyUseCase) GetAll(ctx context.Context) ([]entity.Agency, error) {
	return c.agencyRepo.GetAll()
}

func (c *agencyUseCase) UpdateAgency(ctx context.Context, id uint, p *input.UpdateAgencyRequest) error {
	agency, err := c.agencyRepo.GetByID(id)
	if err != nil {
		return err
	}
	agency.Name = p.Name
	return c.agencyRepo.Update(&agency)
}
func (c *agencyUseCase) DeleteAgency(ctx context.Context, id uint) error {
	agency, err := c.agencyRepo.GetByID(id)
	if err != nil {
		return err
	}
	if agency.ID == 0 {
		return nil // No need to delete if CoQuan does not exist
	}
	return c.agencyRepo.Delete(id)
}
