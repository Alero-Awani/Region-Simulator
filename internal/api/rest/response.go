package rest

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Functions to handle errors

func ErrorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(err.Error())
}

func InternalError(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": msg,
		"data":    data,
	})
}
