package main

import (
	"crucible/vega/handlers"
	"crucible/vega/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	handlers.NewClient("http://localhost", "admin", "password")

	// Routes
	routes.AwxRoutes(app)
	routes.SemaphoreRoutes(app)
	routes.OpsRoutes(app)
	routes.NotFoundRoute(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
