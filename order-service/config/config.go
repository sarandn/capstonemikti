package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v8"
	"github.com/yourusername/order-service/internal/infra/db"
	"github.com/yourusername/order-service/internal/infra/redis"
	"github.com/yourusername/order-service/internal/pkg/utils"
)

type Config struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

func LoadConfig() *Config {
	utils.InfoLogger.Println("Initializing database connection...")
	db := db.InitDB()
	utils.InfoLogger.Println("Database connection initialized")

	utils.InfoLogger.Println("Initializing Redis client...")
	redis := redis.InitRedisClient()
	utils.InfoLogger.Println("Redis client initialized")

	return &Config{
		DB:    db,
		Redis: redis,
	}
}
