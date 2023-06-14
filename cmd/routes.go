package main

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/controllers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// TODO Подумать как сделать controllers.facts.Home
	app.Get("/", controllers.Home)
}
