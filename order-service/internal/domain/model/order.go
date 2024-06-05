package model

import (
    "time"
)

type Order struct {
    OrderID     int       `db:"order_id" json:"order_id"`
    UserID      int       `db:"user_id_fk" json:"user_id_fk"`
    OrderDate   time.Time `db:"order_date" json:"order_date"`
    TotalAmount int       `db:"total_amount" json:"total_amount"`
    CreatedAt   time.Time `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
