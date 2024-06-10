package config

import (
    "log"
    "os"
)

type Config struct {
    ServerAddress string
    DatabaseURL   string
}

func LoadConfig() *Config {
    return &Config{
        ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
        DatabaseURL:   getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/ticketdb"),
    }
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Printf("Peringatan: variabel lingkungan %s tidak disetel, menggunakan nilai default %s", key, defaultValue)
        return defaultValue
    }
    return value
}
