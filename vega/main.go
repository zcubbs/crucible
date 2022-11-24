package main

import (
	"crucible/vega/configs"
	"crucible/vega/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	configs.Bootstrap()
}
func main() {
	app := fiber.New()
	// Routes
	routes.OpsRoutes(app)
	routes.AuthRoutes(app)
	routes.AwxRoutes(app)
	routes.SemaphoreRoutes(app)
	routes.NotFoundRoute(app)

	err := app.Listen(fmt.Sprintf(":%d", configs.Config.API.Port))
	if err != nil {
		log.Fatal(err)
	}
}
