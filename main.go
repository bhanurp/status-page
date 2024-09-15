package main

import (
	"github.com/bhanurp/status-page/app"
	"github.com/bhanurp/status-page/logger"
)

func main() {
	// Create a new instance of the application
	app := &app.App{}
	// Start the application
	if err := app.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
