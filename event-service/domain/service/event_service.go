package service

import (
    "event-service/domain/model"
    "event-service/infra/repository"
)

type EventService struct {
    repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
    return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *model.Event) error {
    return s.repo.CreateEvent(event)
}

func (s *EventService) GetEventByID(eventID uint) (*model.Event, error) {
    return s.repo.GetEventByID(eventID)
}

func (s *EventService) UpdateEvent(event *model.Event) error {
    return s.repo.UpdateEvent(event)
}

func (s *EventService) DeleteEvent(eventID uint) error {
    return s.repo.DeleteEvent(eventID)
}
