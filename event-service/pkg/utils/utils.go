package utils

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}
