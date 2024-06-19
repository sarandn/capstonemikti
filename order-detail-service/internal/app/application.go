package app

import (
	"log"
	"net/http"
	"order-detail-service/config"
	"order-detail-service/internal/domain/service"
	"order-detail-service/internal/infra/repository"
	"order-detail-service/internal/interfaces/api"
	"order-detail-service/pkg/utils"

	"github.com/gorilla/mux"
)

type Application struct {
	Config       *config.Config
	Router       *mux.Router
	OrderDetailRepo    *repository.OrderDetailRepository
	OrderDetailService *service.OrderDetailService
	OrderDetailHandler *api.OrderDetailHandler
}

func NewApplication() *Application {
	cfg := config.LoadConfig()
	orderDetailRepo := repository.NewOrderDetailRepository(cfg.DB)
	orderDetailService := service.NewOrderDetailService(orderDetailRepo)
	orderDetailHandler := api.NewOrderDetailHandler(orderDetailService)

	r := mux.NewRouter()
	app := &Application{
		Config:       cfg,
		Router:       r,
		OrderDetailRepo:    orderDetailRepo,
		OrderDetailService: orderDetailService,
		OrderDetailHandler: orderDetailHandler,
	}

	app.setupRoutes()

	return app
}

func (app *Application) setupRoutes() {
	app.Router.HandleFunc("/ordersdetail", app.OrderDetailHandler.CreateOrderDetail).Methods("POST")
	app.Router.HandleFunc("/ordersdetail/{id}", app.OrderDetailHandler.GetOrderDetail).Methods("GET")
	app.Router.HandleFunc("/getordersdetail", app.OrderDetailHandler.GetOrderDetail).Methods("GET")
	app.Router.HandleFunc("/ordersdetail/{id}", app.OrderDetailHandler.UpdateOrderDetail).Methods("PUT")
	app.Router.HandleFunc("/ordersdetail/{id}", app.OrderDetailHandler.DeleteOrderDetail).Methods("DELETE")
}

func (app *Application) Run(addr string) {
	utils.InfoLogger.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
