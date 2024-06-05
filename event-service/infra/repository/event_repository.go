package repository

import (
    "event-service/domain/model"
    "event-service/infra/db"
    "gorm.io/gorm"
)

type EventRepository struct {
    DB *gorm.DB
}

func NewEventRepository() *EventRepository {
    return &EventRepository{
        DB: db.DB,
    }
}

func (r *EventRepository) CreateEvent(event *model.Event) error {
    return r.DB.Create(event).Error
}

func (r *EventRepository) GetEventByID(eventID uint) (*model.Event, error) {
    var event model.Event
    err := r.DB.First(&event, eventID).Error
    return &event, err
}

func (r *EventRepository) UpdateEvent(event *model.Event) error {
    return r.DB.Save(event).Error
}

func (r *EventRepository) DeleteEvent(eventID uint) error {
    return r.DB.Delete(&model.Event{}, eventID).Error
}
