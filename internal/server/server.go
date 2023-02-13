package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/internal/routes"
)

func NewServer(cfg *config.Config) *fiber.App {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		})
	SetupRoutes(app, cfg.Server.Port)
	return app
}

func SetupRoutes(app *fiber.App, port string) {
	api := app.Group("/api")
	user := api.Group("/user")
	health := api.Group("/health")
	routes.HealthRoutes(health)
	routes.UserRoutes(user)
	err := app.Listen(":" + port)
	if err != nil {
		return
	}
}
