package main

import (
	"order-detail-service/internal/app"

)

func main() {
    application := app.NewApplication()
    application.Run(":8080")
}
