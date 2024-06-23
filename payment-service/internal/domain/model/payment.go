package model

import "time"

type Payment struct {
	PaymentID     int      `json:"payment_id"`
	OrderIDFK     int       `json:"order_id_fk"`
	PaymentDate   time.Time `json:"payment_date"`
	AmountPaid    int       `json:"amount_paid"`
	PaymentMethod string   `json:"payment_method"`
	PaymentStatus string   `json:"payment_status"`
	CreatedAt     string    `json:"created_at"`
}
