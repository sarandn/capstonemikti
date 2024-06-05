package config

import (
	"log"
	"os"

	"github.com/yourusername/go-crud/internal/infra/db"
	"github.com/yourusername/go-crud/internal/infra/redis"
)

type Config struct {
	DB   *sqlx.DB
	Redis *redis.Client
}

func LoadConfig() *Config {
	db := db.InitDB()
	redis := redis.InitRedisClient()

	return &Config{
		DB:   db,
		Redis: redis,
	}
}
