package model

import "time"

type OrderDetail struct {
	OrderDetailID uint `gorm:"primaryKey"`
	OrderID       uint
	TicketID      uint
	Quantity      int
	Subtotal      float64
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
