package db

import (
	"ticket-service/config"
	"ticket-service/domain/model"
)

func Init() {
	db := config.GetDB()
	db.AutoMigrate(&model.Ticket{})
}
