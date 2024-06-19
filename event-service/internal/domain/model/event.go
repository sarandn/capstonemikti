package model

import "time"

type Event struct {
	EventID      int       `json:"event_id" db:"event_id"`
	UserIDFK     int       `json:"user_id_fk" db:"user_id_fk"`
	EventName    string    `json:"event_name" db:"event_name"`
	Image        string    `json:"image" db:"image"`
	Location     string    `json:"location" db:"location"`
	Longitude    float64   `json:"longitude" db:"longitude"`
	Latitude     float64   `json:"latitude" db:"latitude"`
	DateStart    time.Time `json:"date_start" db:"date_start"`
	DateEnd      time.Time `json:"date_end" db:"date_end"`
	Price        int       `json:"price" db:"price"`
	QuantityIDFK int       `json:"quantity_id_fk" db:"quantity_id_fk"`
	CategoryIDFK int       `json:"category_id_fk" db:"category_id_fk"`
	TotalLike    int       `json:"total_like" db:"total_like"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
