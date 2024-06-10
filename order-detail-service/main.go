package main

import (
	"log"
	"net/http"
	"order-detail-service/config"
	"order-detail-service/internal/domain/service"
	"order-detail-service/internal/interfaces/api"
	"order-detail-service/pkg/utils"

	"github.com/gorilla/mux"
)

func main() {
	// Inisialisasi aplikasi
	application := initializeApplication()

	// Jalankan server HTTP
	addr := ":8080" // Ganti dengan alamat port yang diinginkan
	utils.InfoLogger.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, application.Router))
}

func initializeApplication() *app.Application {
	// Inisialisasi konfigurasi
	cfg := config.LoadConfig()

	// Inisialisasi layanan order detail
	orderDetailService := service.NewOrderDetailService(cfg.DB)

	// Inisialisasi handler API order detail
	orderDetailHandler := api.NewOrderDetailHandler(orderDetailService)

	// Inisialisasi router mux
	router := mux.NewRouter()

	// Setup routes
	setupRoutes(router, orderDetailHandler)

	// Return application instance
	return &app.Application{
		Config:             cfg,
		Router:             router,
		OrderDetailService: orderDetailService,
		OrderDetailHandler: orderDetailHandler,
	}
}

func setupRoutes(router *mux.Router, orderDetailHandler *api.OrderDetailHandler) {
	router.HandleFunc("/order-details", orderDetailHandler.CreateOrderDetail).Methods("POST")
	router.HandleFunc("/order-details/{id}", orderDetailHandler.GetOrderDetail).Methods("GET")
	router.HandleFunc("/order-details", orderDetailHandler.GetOrderDetails).Methods("GET")
	router.HandleFunc("/order-details/{id}", orderDetailHandler.UpdateOrderDetail).Methods("PUT")
	router.HandleFunc("/order-details/{id}", orderDetailHandler.DeleteOrderDetail).Methods("DELETE")
}
