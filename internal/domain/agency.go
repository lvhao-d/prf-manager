package domain

type Agency struct {
	ID   string
	Name string
}
type AgencyRepository interface {
	GetAll() ([]Agency, error)
	Create(agency *Agency) error
}
