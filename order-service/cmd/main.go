package main

import (
	"order-service/internal/app"

)

func main() {
    application := app.NewApplication()
    application.Run(":8080")
}
