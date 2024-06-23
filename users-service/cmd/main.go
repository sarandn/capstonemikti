package main

import "users-service/internal/app"

func main() {
	// initialize the aplication
	r := app.InitializedServer()

	r.Logger.Fatal(r.Start(":8006"))
}
