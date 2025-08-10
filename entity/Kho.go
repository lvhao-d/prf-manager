package entity

type Kho struct {
	ID       int
	Ten      string
	CoQuanID int

	// Quan hệ ngược lại (thuộc Cơ quan)
	CoQuan CoQuan `gorm:"foreignKey:CoQuanID" `

	// 1 Kho có nhiều Hồ sơ
	HoSos []HoSo `gorm:"foreignKey:KhoID"`
}
