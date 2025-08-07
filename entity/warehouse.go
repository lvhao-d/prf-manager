package entity

type Kho struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Ten      string `gorm:"type:varchar(255);not null" json:"ten"`
	CoQuanID uint   `gorm:"not null" json:"co_quan_id"`

	// Quan hệ ngược lại (thuộc Cơ quan)
	CoQuan CoQuan `gorm:"foreignKey:CoQuanID" json:"co_quan,omitempty"`

	// 1 Kho có nhiều Hồ sơ
	HoSos []HoSo `gorm:"foreignKey:KhoID" json:"ho_sos,omitempty"`
}
