package api

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "payment-service/internal/domain/model"
    "payment-service/internal/infra/repository"
)

type PaymentHandler struct {
    Repo *repository.PaymentRepository
}

func NewPaymentHandler(repo *repository.PaymentRepository) *PaymentHandler {
    return &PaymentHandler{Repo: repo}
}

func (h *PaymentHandler) CreatePayment(w http.ResponseWriter, r *http.Request) {
    var payment model.Payment
    if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Repo.Create(&payment); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(payment)
}

func (h *PaymentHandler) GetPayment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid payment ID", http.StatusBadRequest)
        return
    }

    payment, err := h.Repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if payment == nil {
        http.Error(w, "Payment not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(payment)
}
