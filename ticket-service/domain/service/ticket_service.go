package service

import (
    "ticket-service/domain/model"
    "ticket-service/infra/repository"
)

type TicketService interface {
    CreateTicket(ticket *model.Ticket) error
    GetTicketByID(id int) (*model.Ticket, error)
    GetAllTickets() ([]*model.Ticket, error)
    UpdateTicket(ticket *model.Ticket) error
    DeleteTicket(id int) error
}

type ticketService struct {
    repo repository.TicketRepository
}

func NewTicketService(repo repository.TicketRepository) TicketService {
    return &ticketService{repo: repo}
}

func (s *ticketService) CreateTicket(ticket *model.Ticket) error {
    return s.repo.Create(ticket)
}

func (s *ticketService) GetTicketByID(id int) (*model.Ticket, error) {
    return s.repo.FindByID(id)
}

func (s *ticketService) GetAllTickets() ([]*model.Ticket, error) {
    return s.repo.FindAll()
}

func (s *ticketService) UpdateTicket(ticket *model.Ticket) error {
    return s.repo.Update(ticket)
}

func (s *ticketService) DeleteTicket(id int) error {
    return s.repo.Delete(id)
}
