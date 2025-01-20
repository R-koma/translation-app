package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func (u *User) GetID() uint {
	return u.ID
}
