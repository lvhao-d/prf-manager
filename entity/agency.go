package entity

type CoQuan struct {
	ID  uint   `gorm:"primaryKey" json:"id"`
	Ten string `gorm:"type:varchar(255);not null" json:"ten"`

	// Quan hệ: 1 Cơ quan có nhiều Kho
	Khos []Kho `gorm:"foreignKey:CoQuanID" json:"khos,omitempty"`

	// Quan hệ: 1 Cơ quan có thể lưu trữ nhiều Hồ sơ (được chuyển đến)
	HoSosLuuTru []HoSo `gorm:"foreignKey:CoQuanLuuTruID" json:"ho_sos_luu_tru,omitempty"`
}
