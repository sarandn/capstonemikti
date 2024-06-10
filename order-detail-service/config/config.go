package config

import (
	"order-detail-service/internal/infra/"
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
