package model

type Category struct {
	CategoryID int    `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
