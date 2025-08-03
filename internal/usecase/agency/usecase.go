package agency

import "prf-manager/internal/domain"

type UseCase struct {
	agencyRepo domain.AgencyRepository
}

func NewUseCase(agencyRepo domain.AgencyRepository) *UseCase {
	return &UseCase{
		agencyRepo: agencyRepo,
	}
}

func (uc *UseCase) GetAllAgencies() ([]domain.Agency, error) {
	return uc.agencyRepo.GetAll()
}
