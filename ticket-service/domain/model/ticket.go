package model

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketID      int    `json:" ticket_id" db:" ticket_id"`
	TicketType    string `json:"ticket_type" db:"ticket_type"`
	TicketPrice   int    `json:"ticket_price" db:"ticket_price"`
	QuantityAvail int    `json:"quantity_avail" db:"quantity_avail"`
}
