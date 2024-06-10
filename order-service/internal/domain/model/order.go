package model

import (
	"time"
)

type Order struct {
	OrderID     int       `gorm:"primaryKey"`
	UserID      int       `gorm:"not null"`
	OrderDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	TotalAmount int
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
