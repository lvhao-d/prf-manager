package domain

type Record struct {
	ID       string
	AgencyID string
	Title    string
}
type RecordRepository interface {
	GetAll() ([]Record, error)
	Create(record *Record) error
}
