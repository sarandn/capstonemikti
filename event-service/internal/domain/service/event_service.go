package service

import (
	"event-service/internal/domain/model"
	"event-service/internal/infra/repository"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *model.Event) error {
	return s.repo.Create(event)
}

func (s *EventService) GetEventByID(eventID int) (*model.Event, error) {
	return s.repo.GetByID(eventID)
}

func (s *EventService) GetAllEvents() ([]*model.Event, error) {
	return s.repo.GetAll()
}

func (s *EventService) UpdateEvent(event *model.Event) error {
	return s.repo.Update(event)
}

func (s *EventService) DeleteEvent(eventID int) error {
	return s.repo.Delete(eventID)
}
