package api

import (
	"encoding/json"
	"net/http"
	"order-detail-service/internal/domain/model"
	"order-detail-service/internal/domain/service"
	"order-detail-service/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderDetailHandler struct {
	Service *service.OrderDetailService
}

func NewOrderDetailHandler(service *service.OrderDetailService) *OrderDetailHandler {
	return &OrderDetailHandler{Service: service}
}

func (h *OrderDetailHandler) CreateOrderDetail(w http.ResponseWriter, r *http.Request) {
	var orderDetail model.OrderDetail
	if err := json.NewDecoder(r.Body).Decode(&orderDetail); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateOrderDetail(&orderDetail); err != nil {
		utils.ErrorLogger.Printf("Failed to create order detail: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	utils.InfoLogger.Println("Order detail created successfully")
	json.NewEncoder(w).Encode(orderDetail)
}

func (h *OrderDetailHandler) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLogger.Printf("Invalid order detail ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderDetail, err := h.Service.GetOrderDetail(int(id))
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get order detail: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(orderDetail)
	utils.InfoLogger.Println("Order detail retrieved successfully")
}

func (h *OrderDetailHandler) GetOrderDetails(w http.ResponseWriter, r *http.Request) {
	orderDetails, err := h.Service.GetOrderDetails()
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get order details: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orderDetails)
	utils.InfoLogger.Println("Order details retrieved successfully")
}

func (h *OrderDetailHandler) UpdateOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLogger.Printf("Invalid order detail ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var orderDetail model.OrderDetail
	if err := json.NewDecoder(r.Body).Decode(&orderDetail); err != nil {
		utils.ErrorLogger.Printf("Failed to decode request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orderDetail.OrderDetailID = uint(id)
	if err := h.Service.UpdateOrderDetail(&orderDetail); err != nil {
		utils.ErrorLogger.Printf("Failed to update order detail: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.InfoLogger.Println("Order detail updated successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orderDetail)
}

func (h *OrderDetailHandler) DeleteOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.ErrorLogger.Printf("Invalid order detail ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteOrderDetail(int(id)); err != nil {
		utils.ErrorLogger.Printf("Failed to delete order detail: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.InfoLogger.Println("Order detail deleted successfully")
	w.WriteHeader(http.StatusNoContent)
}
