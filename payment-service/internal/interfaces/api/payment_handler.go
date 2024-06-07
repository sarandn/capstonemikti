package api

import (
    "net/http"
    "payment-service/internal/domain/model"
    "payment-service/internal/infra/repository"
    "strconv"

    "github.com/gin-gonic/gin"
)

type PaymentHandler struct {
    Repo repository.PaymentRepository
}

func NewPaymentHandler(repo repository.PaymentRepository) *PaymentHandler {
    return &PaymentHandler{Repo: repo}
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
    var payment model.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Repo.Create(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPayment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    payment, err := h.Repo.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) UpdatePayment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var payment model.Payment
    if err := h.Repo.GetByID(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
        return
    }

    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    payment.ID = uint(id)
    if err := h.Repo.Update(&payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.Repo.Delete(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Payment deleted"})
}