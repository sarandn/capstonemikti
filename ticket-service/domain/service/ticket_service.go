package service

import (
	"ticket-service/domain/model"
	"ticket-service/infra/repository"
)

type TicketService struct {
	Repo repository.TicketRepository
}

func (s *TicketService) CreateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	return s.Repo.Create(ticket)
}

func (s *TicketService) GetTickets() ([]model.Ticket, error) {
	return s.Repo.GetAll()
}

func (s *TicketService) GetTicketByID(id uint) (*model.Ticket, error) {
	return s.Repo.GetByID(id)
}

func (s *TicketService) UpdateTicket(ticket *model.Ticket) (*model.Ticket, error) {
	return s.Repo.Update(ticket)
}

func (s *TicketService) DeleteTicket(id uint) error {
	return s.Repo.Delete(id)
}
