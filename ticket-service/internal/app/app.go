package app

import (
	"net/http"
	"os"
	"ticket-service/config"
	"ticket-service/domain/service"
	"ticket-service/infra/db"
	"ticket-service/infra/repository"
	"ticket-service/interfaces"

	"github.com/labstack/echo/v4"
)

func StartApp() {
	db.Init()
	e := echo.New()
	dbInstance := config.GetDB()
	ticketRepo := repository.TicketRepository{DB: dbInstance}          // Menggunakan nilai, bukan pointer
	ticketService := service.TicketService{Repo: &ticketRepo}          // Memberikan pointer ke ticketRepo
	ticketHandler := interfaces.TicketHandler{Service: &ticketService} // Memberikan pointer ke ticketService

	e.POST("/tickets", ticketHandler.CreateTicket)
	e.GET("/tickets", ticketHandler.GetTickets)
	e.GET("/tickets/:id", ticketHandler.GetTicketByID)
	e.PUT("/tickets/:id", ticketHandler.UpdateTicket)
	e.DELETE("/tickets/:id", ticketHandler.DeleteTicket)
	e.POST("/token", ticketHandler.GenerateToken)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
