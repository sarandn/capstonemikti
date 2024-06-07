package interfaces

import (
    "event-service/domain/model"
    "event-service/domain/service"
    "github.com/labstack/echo/v4"
    "net/http"
    "strconv"
)

type EventHandlers struct {
    eventService *service.EventService
}

func NewEventHandlers(eventService *service.EventService) *EventHandlers {
    return &EventHandlers{eventService: eventService}
}

func (h *EventHandlers) CreateEvent(c echo.Context) error {
    var event model.Event
    if err := c.Bind(&event); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    err := h.eventService.CreateEvent(&event)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, event)
}

func (h *EventHandlers) GetEventByID(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    event, err := h.eventService.GetEventByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, event)
}

func (h *EventHandlers) UpdateEvent(c echo.Context) error {
    var event model.Event
    if err := c.Bind(&event); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    err := h.eventService.UpdateEvent(&event)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, event)
}

func (h *EventHandlers) DeleteEvent(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    err := h.eventService.DeleteEvent(uint(id))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusNoContent, nil)
}
