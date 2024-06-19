package repository

import (
	"database/sql"
	"event-service/internal/domain/model"
	"log"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(event *model.Event) error {
	query := `
    INSERT INTO events (user_id, event_name, image, location, longitude, latitude, date_start, date_end, price, quantity_id, category_id, total_like)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING event_id`
	err := r.db.QueryRow(query, event.UserID, event.EventName, event.Image, event.Location, event.Longitude, event.Latitude, event.DateStart, event.DateEnd, event.Price, event.QuantityID, event.CategoryID, event.TotalLike).Scan(&event.EventID)
	if err != nil {
		log.Printf("Failed to create event: %v", err)
		return err
	}
	return nil
}

func (r *EventRepository) GetByID(eventID int) (*model.Event, error) {
	query := `SELECT * FROM events WHERE event_id = $1`
	row := r.db.QueryRow(query, eventID)
	var event model.Event
	err := row.Scan(&event.EventID, &event.UserID, &event.EventName, &event.Image, &event.Location, &event.Longitude, &event.Latitude, &event.DateStart, &event.DateEnd, &event.Price, &event.QuantityID, &event.CategoryID, &event.TotalLike)
	if err != nil {
		log.Printf("Failed to get event: %v", err)
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) GetAll() ([]*model.Event, error) {
	query := `SELECT * FROM events`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Failed to get events: %v", err)
		return nil, err
	}
	defer rows.Close()

	var events []*model.Event
	for rows.Next() {
		var event model.Event
		if err := rows.Scan(&event.EventID, &event.UserID, &event.EventName, &event.Image, &event.Location, &event.Longitude, &event.Latitude, &event.DateStart, &event.DateEnd, &event.Price, &event.QuantityID, &event.CategoryID, &event.TotalLike); err != nil {
			log.Printf("Failed to scan event: %v", err)
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (r *EventRepository) Update(event *model.Event) error {
	query := `
    UPDATE events SET user_id = $1, event_name = $2, image = $3, location = $4, longitude = $5, latitude = $6, date_start = $7, date_end = $8, price = $9, quantity_id = $10, category_id = $11, total_like = $12
    WHERE event_id = $13`
	_, err := r.db.Exec(query, event.UserID, event.EventName, event.Image, event.Location, event.Longitude, event.Latitude, event.DateStart, event.DateEnd, event.Price, event.QuantityID, event.CategoryID, event.TotalLike, event.EventID)
	if err != nil {
		log.Printf("Failed to update event: %v", err)
		return err
	}
	return nil
}

func (r *EventRepository) Delete(eventID int) error {
	query := `DELETE FROM events WHERE event_id = $1`
	_, err := r.db.Exec(query, eventID)
	if err != nil {
		log.Printf("Failed to delete event: %v", err)
		return err
	}
	return nil
}
