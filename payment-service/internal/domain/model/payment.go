package model

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       uint      `gorm:"not null"`
	PaymentDate   time.Time `gorm:"not null"`
	AmountPaid    float64   `gorm:"not null"`
	PaymentMethod string    `gorm:"not null"`
	PaymentStatus string    `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}