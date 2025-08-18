package entity

type Record struct {
	ID              int
	Name            string
	WarehouseID     int
	ArchiveAgencyID int

	PrevWarehouseID     int
	PrevArchiveAgencyID int
	// Quan hệ ngược
	Warehouse *Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
}
