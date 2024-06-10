package service

import (
	"context"
	"order-detail-service/internal/domain/model"
	"order-detail-service/internal/infra/repository"
)

type OrderDetailService struct {
	Repo *repository.OrderDetailRepository
}

func NewOrderDetailService(repo *repository.OrderDetailRepository) *OrderDetailService {
	return &OrderDetailService{Repo: repo}
}

func (s *OrderDetailService) CreateOrderDetail(orderDetail *model.OrderDetail) error {
	return s.Repo.CreateOrderDetail(context.Background(), orderDetail)
}

func (s *OrderDetailService) GetOrderDetail(id uint) (*model.OrderDetail, error) {
	return s.Repo.GetOrderDetail(context.Background(), id)
}

func (s *OrderDetailService) GetOrderDetails() ([]model.OrderDetail, error) {
	return s.Repo.GetOrderDetails(context.Background())
}

func (s *OrderDetailService) UpdateOrderDetail(orderDetail *model.OrderDetail) error {
	return s.Repo.UpdateOrderDetail(context.Background(), orderDetail)
}

func (s *OrderDetailService) DeleteOrderDetail(id uint) error {
	return s.Repo.DeleteOrderDetail(context.Background(), id)
}
