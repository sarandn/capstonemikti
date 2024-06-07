package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/yourusername/order-service/internal/pkg/utils"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		utils.ErrorLogger.Fatalf("Failed to connect to database: %v", err)
	}
	utils.InfoLogger.Println("Database connected successfully")
	return db
}