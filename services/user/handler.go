package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
)

func CreateUser(c *fiber.Ctx) error {
	s := NewUserService()
	user := &database.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	user, err := s.CreateUser(user)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if user == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User already exists",
		})
	}
	return c.JSON(user)
}
