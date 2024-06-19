package model

import "time"

type Event struct {
	EventID    int       `json:"event_id" db:"event_id"`
	UserID     int       `json:"user_id" db:"user_id"`
	EventName  string    `json:"event_name" db:"event_name"`
	Image      string    `json:"image" db:"image"`
	Location   string    `json:"location" db:"location"`
	Longitude  float64   `json:"longitude" db:"longitude"`
	Latitude   float64   `json:"latitude" db:"latitude"`
	DateStart  time.Time `json:"date_start" db:"date_start"`
	DateEnd    time.Time `json:"date_end" db:"date_end"`
	Price      int       `json:"price" db:"price"`
	QuantityID int       `json:"quantity_id" db:"quantity_id"`
	CategoryID int       `json:"category_id" db:"category_id"`
	TotalLike  int       `json:"total_like" db:"total_like"`
}
