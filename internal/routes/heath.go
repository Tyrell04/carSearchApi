package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/services/health"
)

func HealthRoutes(health_router fiber.Router) {
	health_router.Get("/", health.HealthHandler)
}
