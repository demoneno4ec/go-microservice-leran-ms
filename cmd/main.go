package main

import (
	"github.com/demoneno4ec/go-microservice-leran-ms/config"
	"github.com/demoneno4ec/go-microservice-leran-ms/controllers/facts"
	"github.com/demoneno4ec/go-microservice-leran-ms/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	config.SetupRoutes(app)
	app.Use(facts.NotFound)

	app.Listen(":3000")
}
