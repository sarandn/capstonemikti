package interfaces

import (
	"ticket-service/domain/model"
	"ticket-service/domain/service"
	"ticket-service/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	Service *service.TicketService
}

func (h *TicketHandler) CreateTicket(c echo.Context) error {
	ticket := new(model.Ticket)
	if err := c.Bind(ticket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	createdTicket, err := h.Service.CreateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, createdTicket)
}

func (h *TicketHandler) GetTickets(c echo.Context) error {
	tickets, err := h.Service.GetTickets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) GetTicketByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ticket, err := h.Service.GetTicketByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) UpdateTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ticket := new(model.Ticket)
	if err := c.Bind(ticket); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ticket.ID = uint(id)
	updatedTicket, err := h.Service.UpdateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, updatedTicket)
}

func (h *TicketHandler) DeleteTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.Service.DeleteTicket(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Ticket deleted"})
}

func (h *TicketHandler) GenerateToken(c echo.Context) error {
	userID := 1 // Ini adalah contoh. Anda harus mengambil userID dari database.
	token, err := utils.GenerateJWT(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
