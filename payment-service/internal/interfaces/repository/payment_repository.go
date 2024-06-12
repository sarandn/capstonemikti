package repository

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "payment-service/internal/domain/model"

    _ "github.com/lib/pq"
)

type PaymentRepository struct {
    DB *sql.DB
}

func NewPaymentRepository() *PaymentRepository {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) Create(payment *model.Payment) error {
    query := "INSERT INTO payments (amount, status) VALUES ($1, $2) RETURNING id, created_at"
    err := r.DB.QueryRow(query, payment.Amount, payment.Status).Scan(&payment.ID, &payment.CreatedAt)
    return err
}

func (r *PaymentRepository) GetByID(id int) (*model.Payment, error) {
    query := "SELECT id, amount, status, created_at FROM payments WHERE id = $1"
    payment := &model.Payment{}
    err := r.DB.QueryRow(query, id).Scan(&payment.ID, &payment.Amount, &payment.Status, &payment.CreatedAt)
    return payment, err
}
