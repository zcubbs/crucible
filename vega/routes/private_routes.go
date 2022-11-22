package routes

import (
	"crucible/vega/handlers"
	"github.com/gofiber/fiber/v2"
)

// AwxRoutes func for describe group of private routes.
func AwxRoutes(app *fiber.App) {
	route := app.Group("/api/v1/awx")

	route.Get("/ping", handlers.HandlePing)
	route.Get("/runTemplate", handlers.HandleRunTemplate)
	route.Get("/getJobEvents", handlers.HandleGetJobEvents)
}

// SemaphoreRoutes func for describe group of private routes.
func SemaphoreRoutes(app *fiber.App) {
	route := app.Group("/api/v1/semaphore")

	route.Get("/ping", handlers.HandleSemaphorePing)
	route.Post("/auth", handlers.HandleSemaphoreLogin)
	route.Get("/project", handlers.HandleSemaphoreGetProject)
}

func OpsRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/health", handlers.HandleHealthCheck)
}
