package redis

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"github.com/yourusername/order-service/internal/pkg/utils"
)

var rdb *redis.Client

func InitRedisClient() *redis.Client {
	if rdb != nil {
		return rdb
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		utils.ErrorLogger.Fatalf("Failed to connect to Redis: %v", err)
	}

	utils.InfoLogger.Println("Connected to Redis")
	return rdb
}

func GetRedisClient() *redis.Client {
	return rdb
}