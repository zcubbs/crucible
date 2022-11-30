package routes

import (
	"crucible/vega/configs"
	"crucible/vega/handlers"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
)

func AuthRoutes(app *fiber.App) {
	route := app.Group("/auth")

	route.Post("/token", handlers.Login)
}

// AwxRoutes func for describe group of private routes.
func AwxRoutes(app *fiber.App) {
	route := app.Group("/api/awx")

	route.Get("/ping", handlers.HandlePing)
	route.Get("/runTemplate", handlers.HandleRunTemplate)
	route.Get("/getJobEvents", handlers.HandleGetJobEvents)
}

// SemaphoreRoutes func for describe group of private routes.
func SemaphoreRoutes(app *fiber.App) {
	route := app.Group("/api/semaphore")

	route.Use(jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey: []byte(configs.Config.API.TokenSecret),
	}))

	route.Get("/ping", handlers.HandleSemaphorePing)
	route.Post("/auth", handlers.HandleSemaphoreLogin)
	route.Get("/project", handlers.HandleSemaphoreGetProject)
	route.Get("/projects", handlers.HandleSemaphoreGetProjects)
	route.Post("/run-task-template", handlers.HandleSemaphoreRunTaskTemplate)
}

func OpsRoutes(app *fiber.App) {
	route := app.Group("/ops")

	route.Get("/health", handlers.HandleHealthCheck)
}
