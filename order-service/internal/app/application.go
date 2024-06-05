package app

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/yourusername/go-crud/config"
    "github.com/yourusername/go-crud/internal/domain/service"
)

type Application struct {
    Config      *config.Config
    Router      *mux.Router
    OrderService *service.OrderService
}

func NewApplication() *Application {
    cfg := config.LoadConfig()
    orderService := service.NewOrderService(cfg.DB)

    r := mux.NewRouter()
    app := &Application{
        Config:      cfg,
        Router:      r,
        OrderService: orderService,
    }

    app.setupRoutes()

    return app
}

func (app *Application) setupRoutes() {
    app.Router.HandleFunc("/orders", app.createOrder).Methods("POST")
    app.Router.HandleFunc("/orders/{id}", app.getOrder).Methods("GET")
    app.Router.HandleFunc("/orders", app.getOrders).Methods("GET")
    app.Router.HandleFunc("/orders/{id}", app.updateOrder).Methods("PUT")
    app.Router.HandleFunc("/orders/{id}", app.deleteOrder).Methods("DELETE")
}

func (app *Application) Run(addr string) {
    log.Println("Listening on", addr)
    log.Fatal(http.ListenAndServe(addr, app.Router))
}