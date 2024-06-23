package config

import (
	"order-service/internal/infra/repository/db"

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
