package service

import (
    "time"

    "github.com/jmoiron/sqlx"
    "github.com/yourusername/go-crud/internal/domain/model"
)

type OrderService struct {
    DB *sqlx.DB
}

func NewOrderService(db *sqlx.DB) *OrderService {
    return &OrderService{DB: db}
}

func (s *OrderService) CreateOrder(order *model.Order) error {
    order.CreatedAt = time.Now()
    order.UpdatedAt = time.Now()

    return s.DB.QueryRow(`INSERT INTO orders (user_id_fk, order_date, total_amount, created_at, updated_at) 
                          VALUES ($1, $2, $3, $4, $5) RETURNING order_id`,
        order.UserID, order.OrderDate, order.TotalAmount, order.CreatedAt, order.UpdatedAt).Scan(&order.OrderID)
}

func (s *OrderService) GetOrder(id int) (*model.Order, error) {
    var order model.Order
    if err := s.DB.Get(&order, "SELECT * FROM orders WHERE order_id=$1", id); err != nil {
        return nil, err
    }
    return &order, nil
}

func (s *OrderService) GetOrders() ([]model.Order, error) {
    var orders []model.Order
    if err := s.DB.Select(&orders, "SELECT * FROM orders"); err != nil {
        return nil, err
    }
    return orders, nil
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
    order.UpdatedAt = time.Now()
    _, err := s.DB.Exec(`UPDATE orders SET user_id_fk=$1, order_date=$2, total_amount=$3, updated_at=$4 WHERE order_id=$5`,
        order.UserID, order.OrderDate, order.TotalAmount, order.UpdatedAt, order.OrderID)
    return err
}

func (s *OrderService) DeleteOrder(id int) error {
    _, err := s.DB.Exec("DELETE FROM orders WHERE order_id=$1", id)
    return err
}
