package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}
