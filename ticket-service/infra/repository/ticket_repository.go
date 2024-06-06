package repository

import (
    "database/sql"
    "ticket-service/domain/model"
)

type TicketRepository interface {
    Create(ticket *model.Ticket) error
    FindByID(id int) (*model.Ticket, error)
    FindAll() ([]*model.Ticket, error)
    Update(ticket *model.Ticket) error
    Delete(id int) error
}

type ticketRepository struct {
    db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
    return &ticketRepository{db: db}
}

func (r *ticketRepository) Create(ticket *model.Ticket) error {
    _, err := r.db.Exec("INSERT INTO tickets (title, status) VALUES ($1, $2)", ticket.Title, ticket.Status)
    return err
}

func (r *ticketRepository) FindByID(id int) (*model.Ticket, error) {
    row := r.db.QueryRow("SELECT id, title, status FROM tickets WHERE id = $1", id)

    ticket := &model.Ticket{}
    if err := row.Scan(&ticket.ID, &ticket.Title, &ticket.Status); err != nil {
        return nil, err
    }
    return ticket, nil
}

func (r *ticketRepository) FindAll() ([]*model.Ticket, error) {
    rows, err := r.db.Query("SELECT id, title, status FROM tickets")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tickets []*model.Ticket
    for rows.Next() {
        ticket := &model.Ticket{}
        if err := rows.Scan(&ticket.ID, &ticket.Title, &ticket.Status); err != nil {
            return nil, err
        }
        tickets = append(tickets, ticket)
    }
    return tickets, nil
}

func (r *ticketRepository) Update(ticket *model.Ticket) error {
    _, err := r.db.Exec("UPDATE tickets SET title = $1, status = $2 WHERE id = $3", ticket.Title, ticket.Status, ticket.ID)
    return err
}

func (r *ticketRepository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM tickets WHERE id = $1", id)
    return err
}
