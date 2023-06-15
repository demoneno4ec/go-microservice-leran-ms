package config

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/controllers/facts"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/fact", facts.ListFacts)

	app.Post("/fact", facts.Create)
	// Add new route to show single Fact, given `:id`
	app.Get("/fact/:id", facts.Show)
}
