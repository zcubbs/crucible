package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/zcubbs/crucible/vega/configs"
	"github.com/zcubbs/crucible/vega/handlers"
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
	route.Post("/projects", handlers.HandleSemaphoreCreateProject)
	route.Post("/project/:id/repositories", handlers.HandleSemaphoreCreateRepository)
	route.Post("/project/:id/keys", handlers.HandleSemaphoreCreateSSHKey)
	route.Post("/project/:id/inventory", handlers.HandleSemaphoreAddInventory)
	route.Post("/project/:id/environment", handlers.HandleSemaphoreCreateEnvironment)
	route.Post("/project/:id/templates", handlers.HandleSemaphoreCreateTemplate)
	route.Post("/project/:id/tasks", handlers.HandleSemaphoreRunTaskTemplate)
}

func OpsRoutes(app *fiber.App) {
	route := app.Group("/ops")

	route.Get("/health", handlers.HandleHealthCheck)
}
