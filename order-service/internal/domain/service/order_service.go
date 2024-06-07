package service

import (
	"context"
	"order-service/internal/domain/model"
	"order-service/internal/infra/repository"

	"github.com/jmoiron/sqlx"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(db *sqlx.DB) *OrderService {
	return &OrderService{
		repo: repository.NewOrderRepository(db),
	}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.repo.CreateOrder(context.Background(), order)
}

func (s *OrderService) GetOrder(id int) (*model.Order, error) {
	return s.repo.GetOrder(context.Background(), id)
}

func (s *OrderService) GetOrders() ([]model.Order, error) {
	return s.repo.GetOrders(context.Background())
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
	return s.repo.UpdateOrder(context.Background(), order)
}

func (s *OrderService) DeleteOrder(id int) error {
	return s.repo.DeleteOrder(context.Background(), id)
}
