package response

import "github.com/gofiber/fiber/v2"

type success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	response := success{"success", data}

	return c.Status(fiber.StatusOK).JSON(response)
}

func EmptySuccess(c *fiber.Ctx) error {
	response := success{}
	response.Message = "success"
	return c.Status(fiber.StatusOK).JSON(response)
}
