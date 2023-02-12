package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/services/user"
)

func UserRoutes(userRoute fiber.Router) {
	userRoute.Post("/create", user.CreateUser)
}
