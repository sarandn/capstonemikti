package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
