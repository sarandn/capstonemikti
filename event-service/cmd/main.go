package main

import (
    "event-service/internal/app"

)

func main() {


    // Initialize the application
    e := app.Init()

    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}

