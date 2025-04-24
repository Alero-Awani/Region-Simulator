package handlers

import (
	"Region-Simulator/internal/api/rest"
	"Region-Simulator/internal/repository"
	"Region-Simulator/internal/service"
	"github.com/gofiber/fiber/v2"
)

type catalogHandler struct {
	svc service.CatalogService
}

func SetupCatalogRoutes(rh *rest.RestHandler) {
	app := rh.App

	// Create an instance of the catalog service and inject to the handler
	svc := service.CatalogService{
		Repo:   repository.NewCatalogRepository(rh.DB),
		Auth:   rh.Auth,
		Config: rh.Config,
	}

	handler := catalogHandler{
		svc: svc,
	}

	// Public Catalog Endpoints
	app.Get("/products")
	app.Get("/products/:id")
	app.Get("categories")
	app.Get("/categories/:id")

	// Private Catalog Endpoints
	selRoutes := app.Group("/seller")
	// Categories
	selRoutes.Post("/categories", handler.CreateCategories)
	selRoutes.Patch("/categories/:id", handler.EditCategory)
	selRoutes.Delete("/categories/:id", handler.DeleteCategory)

	// Products
	selRoutes.Post("/products", handler.CreateProducts)
	selRoutes.Get("/products", handler.GetProducts)
	selRoutes.Get("/products/:id", handler.GetProduct)
	selRoutes.Put("/products/:id", handler.EditProducts)
	selRoutes.Patch("/products/:id", handler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", handler.DeleteProduct)
}

func (h *catalogHandler) CreateCategories(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "category endpoint", nil)
}

func (h *catalogHandler) EditCategory(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "edit category endpoint", nil)
}

func (h *catalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "delete category endpoint", nil)
}

func (h *catalogHandler) CreateProducts(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "create product endpoint", nil)
}

func (h *catalogHandler) EditProducts(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "edit product endpoint", nil)
}

func (h *catalogHandler) GetProducts(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "get product endpoint", nil)
}

func (h *catalogHandler) GetProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "get product by ID", nil)
}

func (h *catalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "update stock endpoint", nil)
}

func (h *catalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	return rest.SuccessResponse(ctx, "delete product endpoint", nil)

}
