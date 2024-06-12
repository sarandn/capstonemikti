package app

import (
	"log"
	"net/http"
	"order-service/config"
	"order-service/internal/domain/service"
	"order-service/internal/infra/repository"
	"order-service/internal/interfaces/api"
	"order-service/pkg/utils"

	"github.com/gorilla/mux"
)

type Application struct {
	Config       *config.Config
	Router       *mux.Router
	OrderRepo    *repository.OrderRepository
	OrderService *service.OrderService
	OrderHandler *api.OrderHandler
}

func NewApplication() *Application {
	cfg := config.LoadConfig()
	orderRepo := repository.NewOrderRepository(cfg.DB)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := api.NewOrderHandler(orderService)

	r := mux.NewRouter()
	app := &Application{
		Config:       cfg,
		Router:       r,
		OrderRepo:    orderRepo,
		OrderService: orderService,
		OrderHandler: orderHandler,
	}

	app.setupRoutes()

	return app
}

func (app *Application) setupRoutes() {
	app.Router.HandleFunc("/orders", app.OrderHandler.CreateOrder).Methods("POST")
	app.Router.HandleFunc("/orders/{id}", app.OrderHandler.GetOrder).Methods("GET")
	app.Router.HandleFunc("/orders", app.OrderHandler.GetOrders).Methods("GET")
	app.Router.HandleFunc("/orders/{id}", app.OrderHandler.UpdateOrder).Methods("PUT")
	app.Router.HandleFunc("/orders/{id}", app.OrderHandler.DeleteOrder).Methods("DELETE")
}

func (app *Application) Run(addr string) {
	utils.InfoLogger.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
