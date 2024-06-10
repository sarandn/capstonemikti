package app

import (
    "ticket-service/config"
    "ticket-service/infra/db"
    "ticket-service/interfaces"

    "github.com/labstack/echo/v4"
)

func Run() {
    cfg := config.LoadConfig()
    database := db.Connect(cfg)
    e := echo.New()
    interfaces.RegisterHandlers(e, database)
    e.Logger.Fatal(e.Start(cfg.ServerAddress))
}
