package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewUserController(userService *services.UserService, config config.Impl) *UserController {
	return &UserController{UserService: *userService, Impl: config}
}

type UserController struct {
	services.UserService
	config.Impl
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/user", controller.Create)
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request models.LoginRequest
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	controller.UserService.Create(request.Username, request.Password, request.Roles)
	return c.JSON(fiber.Map{"message": "User created successfully!"})
}
