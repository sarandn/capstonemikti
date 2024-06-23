package service

import (
	"context"
	"order-service/internal/domain/model"
	"order-service/internal/infra/repository"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{Repo: repo}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.Repo.CreateOrder(context.Background(), order)
}

func (s *OrderService) GetOrder(id int) (*model.Order, error) {
	return s.Repo.GetOrder(context.Background(), id)
}

func (s *OrderService) GetOrders() ([]model.Order, error) {
	return s.Repo.GetOrders(context.Background())
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
	return s.Repo.UpdateOrder(context.Background(), order)
}

func (s *OrderService) DeleteOrder(id int) error {
	return s.Repo.DeleteOrder(context.Background(), id)
}