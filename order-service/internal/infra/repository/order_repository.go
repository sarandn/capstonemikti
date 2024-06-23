package repository

import (
	"context"
	"order-service/internal/domain/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	return r.DB.WithContext(ctx).Create(order).Error
}

func (r *OrderRepository) GetOrder(ctx context.Context, id int) (*model.Order, error) {
	var order model.Order
	if err := r.DB.WithContext(ctx).First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	if err := r.DB.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	return r.DB.WithContext(ctx).Save(order).Error
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&model.Order{}, id).Error
}
