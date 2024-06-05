package repository

import (
	"payment-service/internal/domain/model"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *model.Payment) error
	GetByID(id uint) (*model.Payment, error)
	Update(payment *model.Payment) error
	Delete(id uint) error
}

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{DB: db}
}

func (r *paymentRepository) Create(payment *model.Payment) error {
	return r.DB.Create(payment).Error
}

func (r *paymentRepository) GetByID(id uint) (*model.Payment, error) {
	var payment model.Payment
	if err := r.DB.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(payment *model.Payment) error {
	return r.DB.Save(payment).Error
}

func (r *paymentRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Payment{}, id).Error
}