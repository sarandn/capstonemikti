package db

import (
    "log"
    "os"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "order-service/internal/pkg/utils"
)

func InitDB() *sqlx.DB {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        utils.ErrorLogger.Fatalf("Failed to connect to database: %v", err)
    }
    utils.InfoLogger.Println("Database connected successfully")
    return db
}
