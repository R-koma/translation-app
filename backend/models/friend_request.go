package models

import (
	"time"

	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey"`
	SenderID   uint   `gorm:"not null"`
	ReceiverID uint   `gorm:"not null"`
	Status     string `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
