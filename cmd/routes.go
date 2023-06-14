package main

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/controllers/facts"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/fact", facts.ListFacts)

	app.Post("/fact", facts.Create)
}
