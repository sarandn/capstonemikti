package db

import (
	"database/sql"
	"log"
	"ticket-service/config"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Tidak dapat terhubung ke database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Tidak dapat melakukan ping ke database: %v", err)
	}

	return db
}
