package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/yourusername/go-crud/internal/domain/model"
	"github.com/yourusername/go-crud/internal/pkg/utils"
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
	if err != nil {
		utils.ErrorLogger.Printf("Failed to create order: %v", err)
		return err
	}
	utils.InfoLogger.Println("Order created successfully")
	return nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, id int) (*model.Order, error) {
	var order model.Order
	err := r.DB.GetContext(ctx, &order, "SELECT * FROM orders WHERE order_id=$1", id)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get order: %v", err)
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	err := r.DB.SelectContext(ctx, &orders, "SELECT * FROM orders")
	if err != nil {
		utils.ErrorLogger.Printf("Failed to get orders: %v", err)
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE orders SET user_id_fk=$1, order_date=$2, total_amount=$3, updated_at=$4 WHERE order_id=$5`,
		order.UserID, order.OrderDate, order.TotalAmount, order.UpdatedAt, order.OrderID)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to update order: %v", err)
		return err
	}
	utils.InfoLogger.Println("Order updated successfully")
	return nil
}

func (r *OrderRepository) DeleteOrder(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM orders WHERE order_id=$1", id)
	if err != nil {
		utils.ErrorLogger.Printf("Failed to delete order: %v", err)
		return err
	}
	utils.InfoLogger.Println("Order deleted successfully")
	return nil
}
