package main

import (
    "log"
    "net/http"
    "payment-service/config"
    "payment-service/internal/app"
)

func main() {
    config.LoadConfig()

    app := app.NewApp()
    if err := app.Run(); err != nil {
        log.Fatalf("could not run app: %v", err)
    }
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", app.Router))
}