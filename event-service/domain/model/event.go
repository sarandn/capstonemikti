package model

import "gorm.io/gorm"

type Event struct {
    gorm.Model
    UserID     uint   `json:"user_id"`
    EventName  string `json:"event_name"`
    Image      string `json:"image"`
    Location   string `json:"location"`
    Longitude  float64 `json:"longitude"`
    Latitude   float64 `json:"latitude"`
    DateStart  string `json:"date_start"`
    DateEnd    string `json:"date_end"`
    Price      int    `json:"price"`
    QuantityID uint   `json:"quantity_id"`
    CategoryID uint   `json:"category_id"`
    TotalLike  int    `json:"total_like"`
}
