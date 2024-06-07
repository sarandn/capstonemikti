package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourusername/go-crud/internal/domain/model"
	"github.com/yourusername/go-crud/internal/domain/service"
	"github.com/yourusername/go-crud/internal/pkg/utils"
)

type OrderHandler struct {
	Service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateOrder(&order); err != nil {
		utils.ErrorLogger.Printf("Failed to create order: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.InfoLogger.Println("Order created successfully")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLogger.Printf("Invalid order ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.Service.GetOrder(id)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get order: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
	utils.InfoLogger.Println("Order retrieved successfully")
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.Service.GetOrders()
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get orders: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
	utils.InfoLogger.Println("Orders retrieved successfully")
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLogger.Printf("Invalid order ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body:
