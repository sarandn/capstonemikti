package db

import (
	"log"
	"os"
	"users-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB {
	config.LoadConfig()

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbDatabase + " port=" + dbPort + " sslmode=" + dbSSLMode + " TimeZone=" + dbTimeZone

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
