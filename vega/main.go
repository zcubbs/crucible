package main

import (
	"crucible/core/postgres"
	"crucible/vega/configs"
	"crucible/vega/queries"
	"crucible/vega/routes"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	configs.Bootstrap()
}
func main() {
	db := postgres.Connect(postgres.DBConfig{
		Host:     configs.Config.Postgres.Host,
		Port:     configs.Config.Postgres.Port,
		User:     configs.Config.Postgres.Username,
		Password: configs.Config.Postgres.Password,
		Database: configs.Config.Postgres.Database,
		SslMode:  configs.Config.Postgres.SslMode,
		Verbose:  configs.Config.Postgres.Verbose,
	})

	queries.Database = &queries.DB{DB: db}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "Vega v1.0.0",
	})
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${ip}]:${port} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Europe/Paris",
	}))

	// Routes
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Vega Metrics"}))
	routes.OpsRoutes(app)
	routes.AuthRoutes(app)
	routes.AwxRoutes(app)
	routes.SemaphoreRoutes(app)
	routes.NotFoundRoute(app)

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// Start server
	log.Infof("Starting vega on port %d", configs.Config.API.Port)
	cover := figure.NewColorFigure("VEGA", "doom", "green", true)
	cover.Print()
	fmt.Print("v1.0.0 by @zcubbs\n\n")
	log.Fatal(app.Listen(fmt.Sprintf(":%d", configs.Config.API.Port)))
}
