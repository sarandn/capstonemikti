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
	ticketRepo := repository.TicketRepository{DB: dbInstance}          // Menggunakan nilai, bukan pointer
	ticketService := service.TicketService{Repo: ticketRepo}           // Menggunakan nilai, bukan pointer
	ticketHandler := interfaces.TicketHandler{Service: &ticketService} // Memberikan pointer ke ticketService

	e.POST("/tickets", ticketHandler.CreateTicket)
	e.GET("/tickets", ticketHandler.GetTickets)
	e.GET("/tickets/:id", ticketHandler.GetTicketByID)
	e.PUT("/tickets/:id", ticketHandler.UpdateTicket)
	e.DELETE("/tickets/:id", ticketHandler.DeleteTicket)
	e.POST("/token", ticketHandler.GenerateToken)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	e.Logger.Fatal(e.Start(":" + port))

	return e
}
