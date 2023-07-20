package models

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncreament" json:"id"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
