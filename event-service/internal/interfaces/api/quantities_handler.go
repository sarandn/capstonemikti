package api

import (
	"encoding/json"
	"event-service/internal/domain/model"
	"event-service/internal/domain/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QuantitiesHandler struct {
	service *service.QuantitiesService
}

func NewQuantitiesHandler(service *service.QuantitiesService) *QuantitiesHandler {
	return &QuantitiesHandler{service: service}
}

func (h *QuantitiesHandler) CreateQuantities(w http.ResponseWriter, r *http.Request) {
	var quantity model.Quantities
	if err := json.NewDecoder(r.Body).Decode(&quantity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateQuantities(&quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quantity)
}

func (h *QuantitiesHandler) GetQuantitiesByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	quantity, err := h.service.GetQuantitiesByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quantity)
}

func (h *QuantitiesHandler) GetAllQuantities(w http.ResponseWriter, r *http.Request) {
	quantities, err := h.service.GetAllQuantities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quantities)
}

func (h *QuantitiesHandler) UpdateQuantities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var quantity model.Quantities
	if err := json.NewDecoder(r.Body).Decode(&quantity); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	quantity.QuantityID = id
	if err := h.service.UpdateQuantities(&quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quantity)
}

func (h *QuantitiesHandler) DeleteQuantities(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteQuantities(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
