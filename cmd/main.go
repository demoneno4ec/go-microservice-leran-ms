package main

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Div Rhino!")
	})

	app.Listen(":3000")
}
