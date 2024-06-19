package model

type Category struct {
	CategoryIDFK   int    `json:"category_id_fk" db:"category_id_fk"`
	CategoryName string `json:"category_name" db:"category_name"`
}
