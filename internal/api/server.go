package api

import (
	"Region-Simulator/Config"
	"Region-Simulator/internal/api/rest"
	"Region-Simulator/internal/api/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config Config.AppConfig) {
	app := fiber.New()
	rh := &rest.RestHandler{
		App: app,
	}
	setupRoutes(rh)
	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	//user handler
	handlers.SetupUserRoutes(rh)

	// transactions
	//	catalog
}
