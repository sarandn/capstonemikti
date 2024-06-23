package model

import (
	"time"
)

type Order struct {
	OrderID     int       `json:"order_id"`
	UserIDFK    int       `json:"user_id_fk"`
	OrderDate   time.Time `json:"order_date"`
	TotalAmount int       `json:"total_amount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
