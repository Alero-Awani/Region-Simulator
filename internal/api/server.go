package api

import (
	"Region-Simulator/config"
	"Region-Simulator/internal/api/rest"
	"Region-Simulator/internal/api/rest/handlers"
	"Region-Simulator/internal/domain"
	"Region-Simulator/internal/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Address{},
		&domain.BankAccount{},
		&domain.Category{},
		&domain.Product{},
		&domain.Cart{},
	)
	if err != nil {
		log.Fatalf("error on running the migration: %v\n", err)
	}

	log.Println("Migration was successful")

	// cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3030",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS, PUT, DELETE, PATCH",
	})
	app.Use(c)

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
	handlers.SetupCatalogRoutes(rh)
}
