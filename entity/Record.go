package entity

type Record struct {
	ID              int
	Name            string
	WarehouseID     int
	ArchiveAgencyID string

	// Quan hệ ngược
	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
}
