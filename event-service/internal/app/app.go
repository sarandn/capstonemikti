package app

import (
    "event-service/config"
    "event-service/domain/service"

    "event-service/infra/repository"
    "event-service/interfaces"
    "github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
    // Load configuration
    config.LoadConfig()

    // Connect to the database

    // Initialize repositories
    eventRepo := repository.NewEventRepository()

    // Initialize services
    eventService := service.NewEventService(eventRepo)

    // Initialize handlers
    eventHandlers := interfaces.NewEventHandlers(eventService)

    // Create Echo instance
    e := echo.New()

    // Define routes
    e.POST("/events", eventHandlers.CreateEvent)
    e.GET("/events/:id", eventHandlers.GetEventByID)
    e.PUT("/events/:id", eventHandlers.UpdateEvent)
    e.DELETE("/events/:id", eventHandlers.DeleteEvent)

    return e
}
