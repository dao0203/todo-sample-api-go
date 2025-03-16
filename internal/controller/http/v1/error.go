package v1

import (
	"github.com/gofiber/fiber/v3"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func errorResponse(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.Status(code).JSON(response{msg})
}
