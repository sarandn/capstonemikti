package interfaces

import (
	"net/http"
	"strconv"
	"ticket-service/domain/model"
	"ticket-service/domain/service"
	"ticket-service/pkg/utils"

	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	Service *service.TicketService
}

// CreateTicket menangani pembuatan tiket baru
func (h *TicketHandler) CreateTicket(c echo.Context) error {
	ticket := new(model.Ticket)
	if err := c.Bind(ticket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Payload permintaan tidak valid"})
	}
	createdTicket, err := h.Service.CreateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, createdTicket)
}

// GetTicket menangani pengambilan semua tiket
func (h *TicketHandler) GetTicket(c echo.Context) error {
	ticket, err := h.Service.GetTicket()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ticket)
}

// GetTicketByID menangani pengambilan tiket berdasarkan ID
func (h *TicketHandler) GetTicketByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID tiket tidak valid"})
	}
	ticket, err := h.Service.GetTicketByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, ticket)
}

// UpdateTicket menangani pembaruan tiket yang sudah ada
func (h *TicketHandler) UpdateTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID tiket tidak valid"})
	}
	ticket := new(model.Ticket)
	if err := c.Bind(ticket); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Payload permintaan tidak valid"})
	}
	ticket.TicketID = int(id)
	updatedTicket, err := h.Service.UpdateTicket(ticket)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedTicket)
}

// DeleteTicket menangani penghapusan tiket berdasarkan ID
func (h *TicketHandler) DeleteTicket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID tiket tidak valid"})
	}
	if err := h.Service.DeleteTicket(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tiket berhasil dihapus"})
}

// GenerateToken menghasilkan token JWT untuk pengguna
func (h *TicketHandler) GenerateToken(c echo.Context) error {
	userID := 1
	token, err := utils.GenerateJWT(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
