package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	// load .env
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbDatabase + " port=" + dbPort + " sslmode=" + dbSSLMode + " TimeZone=" + dbTimeZone

	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
	return db
}
