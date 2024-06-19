package model

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketType    string `json:"ticket_type"`
	TicketPrice   int    `json:"ticket_price"`
	QuantityAvail int    `json:"quantity_avail"`
	Title         string `json:"title,omitempty"`
	Description   string `json:"description,omitempty"`
	Status        string `json:"status,omitempty"`
}
