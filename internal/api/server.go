package api

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/api/rest"
	"Region-Simulator/internal/api/rest/handlers"
	"Region-Simulator/internal/domain"
	"Region-Simulator/internal/helper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	log.Println("Database Connected")

	// run migration
	err = db.AutoMigrate(&domain.User{}, &domain.BankAccount{})
	if err != nil {
		log.Fatalf("error on running the migration: %v\n", err)
	}

	log.Println("Migration was successful")

	auth := helper.SetupAuth(config.AppSecret)

	rh := &rest.RestHandler{
		App:    app,
		DB:     db,
		Auth:   auth,
		Config: config,
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
