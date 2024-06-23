package repository

import (
	"ticket-service/domain/model"

	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func (r *TicketRepository) Create(ticket *model.Ticket) (*model.Ticket, error) {
	if err := r.DB.Create(ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *TicketRepository) GetAll() ([]model.Ticket, error) {
	var ticket []model.Ticket
	if err := r.DB.Find(&ticket).Error; err != nil {
		return nil, err
	}
	return ticket, nil
}

func (r *TicketRepository) GetByID(id uint) (*model.Ticket, error) {
	var ticket model.Ticket
	if err := r.DB.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepository) Update(ticket *model.Ticket) (*model.Ticket, error) {
	result := r.DB.Model(&model.Ticket{}).Where("ticket_id = ?", ticket.TicketID).Updates(map[string]interface{}{
		"event_id_fk":    ticket.EventIDFK,
		"ticket_type":    ticket.TicketType,
		"ticket_price":   ticket.TicketPrice,
		"quantity_avail": ticket.QuantityAvail,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return ticket, nil
}

func (r *TicketRepository) Delete(id uint) error {
	if err := r.DB.Delete(&model.Ticket{}, id).Error; err != nil {
		return err
	}
	return nil
}
