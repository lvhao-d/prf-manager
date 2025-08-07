package entity

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}
