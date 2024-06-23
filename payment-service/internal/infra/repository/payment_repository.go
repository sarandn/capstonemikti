package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"payment-service/config"
	"payment-service/internal/domain/model"

	_ "github.com/lib/pq"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository() *PaymentRepository {
	config.LoadConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return &PaymentRepository{DB: db}
}

func (r *PaymentRepository) Create(payment *model.Payment) error {
	query := "INSERT INTO payments (payment_id, order_id, payment_date, payment_method, amount_paid, payment_status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at"
	err := r.DB.QueryRow(query, payment.PaymentID, payment.OrderIDFK, payment.PaymentDate, payment.PaymentMethod, payment.AmountPaid, payment.PaymentStatus).Scan(&payment.PaymentID, &payment.CreatedAt)
	return err
}

func (r *PaymentRepository) GetByID(id int) (*model.Payment, error) {
	query := "SELECT id, amount, status, created_at FROM payments WHERE id = $1"
	payment := &model.Payment{}
	err := r.DB.QueryRow(query, id).Scan(&payment.PaymentID, payment.OrderIDFK, &payment.PaymentDate, &payment.PaymentMethod,&payment.AmountPaid, &payment.PaymentStatus, &payment.CreatedAt)
	return payment, err
}
