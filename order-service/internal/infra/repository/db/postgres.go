package db

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/joho/godotenv/autoload"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() *sqlx.DB {
    if db != nil {
        return db
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    dsn := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=" + dbSSLMode
    conn, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalln(err)
    }

    db = conn
    return db
}

func GetDB() *sqlx.DB {
    return db
}
