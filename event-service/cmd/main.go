package main

import (
	"event-service/config"
	"event-service/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	app.Start(cfg)
}
