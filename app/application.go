package app

import "os"

type App struct {
	token string
}

func (a *App) Run() error {
	// Initialize the application
	a.initialize()
	return nil
}

func (a *App) initialize() {
	// Initialize the configuration
	a.initializeConfig()
}

func (a *App) initializeConfig() {
	// Read the configuration from the .env file
	a.token = os.Getenv("STATUS_PAGE_BEARER_TOKEN")
}
