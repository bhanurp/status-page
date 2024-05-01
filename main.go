package main

import (
	"log"
	"status-page/app"
)

func main() {
	// Create a new instance of the application
	app := &app.App{}
	// Start the application
	if err := app.Run(); err != nil {
		log.Fatalf(err.Error())
	}
}
