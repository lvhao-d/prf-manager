package entity

type CoQuan struct {
	ID  int
	Ten string

	// Quan hệ: 1 Cơ quan có nhiều Kho
	Khos []Kho `gorm:"foreignKey:CoQuanID" `
}
