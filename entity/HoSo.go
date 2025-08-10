package entity

type HoSo struct {
	ID             int
	Ten            string
	KhoID          int
	CoQuanLuuTruID string

	// Quan hệ ngược
	Kho Kho `gorm:"foreignKey:KhoID"`
}
