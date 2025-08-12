package entity

type Warehouse struct {
	ID       int
	Name     string
	AgencyID int

	// Quan hệ ngược lại (thuộc Cơ quan)
	Agency *Agency `gorm:"foreignKey:AgencyID" json:"agency,omitempty" `

	// 1 Kho có nhiều Hồ sơ
	Records []Record `gorm:"foreignKey:WarehouseID"`
}
