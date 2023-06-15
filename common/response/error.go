package response

import (
	"github.com/gofiber/fiber/v2"
)

type errorStructure struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(c *fiber.Ctx, code int, message string) error {
	response := errorStructure{"error", message}

	return c.Status(code).JSON(response)
}

func NotFound(c *fiber.Ctx) error {
	return Error(c, fiber.StatusNotFound, "Url не задан")
}
