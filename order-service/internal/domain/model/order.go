package model

import (
	"time"
)

type Order struct {
	OrderID     int       `gorm:"column:order_id;primaryKey;autoIncrement"`
	UserIDFK    int       `gorm:"column:user_id_fk"`
	OrderDate   time.Time `gorm:"column:order_date"`
	TotalAmount int       `gorm:"column:total_amount"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
