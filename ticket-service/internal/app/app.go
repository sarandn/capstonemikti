package app

import (
	"os"
	"ticket-service/config"
	"ticket-service/domain/service"
	"ticket-service/infra/db"
	"ticket-service/infra/repository"
	"ticket-service/interfaces"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	db.Init()
	e := echo.New()
	dbInstance := config.GetDB()
	ticketRepo := repository.TicketRepository{DB: dbInstance}
	ticketService := service.TicketService{Repo: ticketRepo}
	ticketHandler := interfaces.TicketHandler{Service: &ticketService}

	e.POST("/ticket", ticketHandler.CreateTicket)
	e.GET("/ticket", ticketHandler.GetTicket)
	e.GET("/ticket/:id", ticketHandler.GetTicketByID)
	e.PUT("/ticket/:id", ticketHandler.UpdateTicket)
	e.DELETE("/ticket/:id", ticketHandler.DeleteTicket)
	e.POST("/token", ticketHandler.GenerateToken)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8005"
	}
	e.Logger.Fatal(e.Start(":" + port))
	return e
}
