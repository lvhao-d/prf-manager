package entity

type HoSo struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Ten            string `gorm:"type:varchar(255);not null" json:"ten"`
	KhoID          uint   `gorm:"not null" json:"kho_id"`
	CoQuanLuuTruID *uint  `json:"co_quan_luu_tru_id"` // nullable

	// Quan hệ ngược
	Kho Kho `gorm:"foreignKey:KhoID" json:"kho,omitempty"`

	// Nếu hồ sơ đã chuyển lưu trữ, liên kết tới Cơ quan lưu trữ
	CoQuanLuuTru *CoQuan `gorm:"foreignKey:CoQuanLuuTruID" json:"co_quan_luu_tru,omitempty"`
}
