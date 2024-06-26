package config

import (
	"order-detail-service/internal/infra/db"

	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func LoadConfig() *Config {
	database := db.InitDB()
	return &Config{
		DB: database,
	}
}
