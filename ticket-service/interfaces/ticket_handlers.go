package interfaces

import (
    "database/sql"
    "net/http"
    "strconv"
    "ticket-service/domain/model"
    "ticket-service/domain/service"
    "ticket-service/infra/repository"

    "github.com/labstack/echo/v4"
)

type TicketHandler struct {
    service service.TicketService
}

func NewTicketHandler(service service.TicketService) *TicketHandler {
    return &TicketHandler{service: service}
}

func RegisterHandlers(e *echo.Echo, db *sql.DB) {
    repo := repository.NewTicketRepository(db)
    svc := service.NewTicketService(repo)
    handler := NewTicketHandler(svc)

    e.POST("/tickets", handler.CreateTicket)
    e.GET("/tickets/:id", handler.GetTicketByID)
    e.GET("/tickets", handler.GetAllTickets)
    e.PUT("/tickets/:id", handler.UpdateTicket)
    e.DELETE("/tickets/:id", handler.DeleteTicket)
}

func (h *TicketHandler) CreateTicket(c echo.Context) error {
    var ticket model.Ticket
    if err := c.Bind(&ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if err := h.service.CreateTicket(&ticket); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, ticket)
}

func (h *TicketHandler) GetTicketByID(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    ticket, err := h.service.GetTicketByID(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) GetAllTickets(c echo.Context) error {
    tickets, err := h.service.GetAllTickets()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) UpdateTicket(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    var ticket model.Ticket
    if err := c.Bind(&ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    ticket.ID = id
    if err := h.service.UpdateTicket(&ticket); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) DeleteTicket(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if err := h.service.DeleteTicket(id); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
