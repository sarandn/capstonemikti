package model

type Quantities struct {
	QuantityID       int `json:"quantity_id" db:"quantity_id"`
	PurchaseQuantity int `json:"purchase_quantity" db:"purchase_quantity"`
}
