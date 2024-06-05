package main

import "github.com/yourusername/go-crud/internal/app"

func main() {
    application := app.NewApplication()
    application.Run(":8080")
}
