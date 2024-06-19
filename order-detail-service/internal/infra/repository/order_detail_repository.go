package repository

import (
	"context"
	"order-detail-service/internal/domain/model"
	

	"gorm.io/gorm"
)

type OrderDetailRepository struct {
	DB *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{DB: db}
}

func (r *OrderDetailRepository) CreateOrderDetail(ctx context.Context, orderDetail *model.OrderDetail) error {
	return r.DB.WithContext(ctx).Create(orderDetail).Error
}

func (r *OrderDetailRepository) GetOrderDetail(ctx context.Context, id int) (*model.OrderDetail, error) {
	var orderDetail model.OrderDetail
	if err := r.DB.WithContext(ctx).First(&orderDetail, id).Error; err != nil {
		return nil, err
	}
	return &orderDetail, nil
}

func (r *OrderDetailRepository) GetOrderDetails(ctx context.Context) ([]model.OrderDetail, error) {
	var orderDetail []model.OrderDetail
	if err := r.DB.WithContext(ctx).Find(&orderDetail).Error; err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func (r *OrderDetailRepository) UpdateOrderDetail(ctx context.Context, orderDetail *model.OrderDetail) error {
	return r.DB.WithContext(ctx).Save(orderDetail).Error
}

func (r *OrderDetailRepository) DeleteOrderDetail(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&model.OrderDetail{}, id).Error
}
