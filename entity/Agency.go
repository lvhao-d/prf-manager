package entity

type Agency struct {
	ID   int
	Name string

	// Quan hệ: 1 Cơ quan có nhiều Kho
	Warehouses []Warehouse `gorm:"foreignKey:AgencyID" `
}
