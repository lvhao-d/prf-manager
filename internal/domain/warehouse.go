package domain

type Warehouse struct {
	ID       string
	Location string
}
type WarehouseRepository interface {
	GetAll() ([]Warehouse, error)
	Create(warehouse *Warehouse) error
}
