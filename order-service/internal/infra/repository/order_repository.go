package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"order-service/internal/domain/model"
)

type OrderRepository struct {
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	_, err := r.DB.ExecContext(ctx, `INSERT INTO orders (user_id_fk, order_date, total_amount, created_at, updated_at) 
									  VALUES ($1, $2, $3, $4, $5)`,
		order.UserID, order.OrderDate, order.TotalAmount, order.CreatedAt, order.UpdatedAt)
	return err
}

func (r *OrderRepository) GetOrder(ctx context.Context, id int) (*model.Order, error) {
	var order model.Order
	err := r.DB.GetContext(ctx, &order, "SELECT * FROM orders WHERE order_id=$1", id)
	return &order, err
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	err := r.DB.SelectContext(ctx, &orders, "SELECT * FROM orders")
	return orders, err
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE orders SET user_id_fk=$1, order_date=$2, total_amount=$3, updated_at=$4 WHERE order_id=$5`,
		order.UserID, order.OrderDate, order.TotalAmount, order.UpdatedAt, order.OrderID)
	return err
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM orders WHERE order_id=$1", id)
	return err
}
