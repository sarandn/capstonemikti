package app

import (
    "payment-service/internal/infra/repository"
    "payment-service/internal/interfaces/api"
    "github.com/gorilla/mux"
    "log"
)

type App struct {
    Router *mux.Router
}

func NewApp() *App {
    app := &App{
        Router: mux.NewRouter(),
    }
    repo := repository.NewPaymentRepository()
    handler := api.NewPaymentHandler(repo)
    app.setRoutes(handler)
    return app
}

func (a *App) setRoutes(handler *api.PaymentHandler) {
    a.Router.HandleFunc("/payments", handler.CreatePayment).Methods("POST")
    a.Router.HandleFunc("/payments/{id:[0-9]+}", handler.GetPayment).Methods("GET")
    a.Router.HandleFunc("/payments/{id:[0-9]+}", handler.GetPayment).Methods("DELETE")
}

func (a *App) Run() error {
    log.Println("Starting the application...")
    return nil
}