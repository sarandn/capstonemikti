package model

type Ticket struct {
	TicketID      int    `json:"ticket_id" db:"ticket_id"`
	EventIDFK     int    `json:"event_id_fk" db:"event_id_fk"`
	TicketType    string `json:"ticket_type" db:"ticket_type"`
	TicketPrice   int    `json:"ticket_price" db:"ticket_price"`
	QuantityAvail int    `json:"quantity_avail" db:"quantity_avail"`
	
}

type Table interface {
	TableName() string
}

func (Ticket) TableName() string {
	return "ticket"
}
