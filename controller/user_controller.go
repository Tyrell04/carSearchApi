package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/common"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/middleware"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewUserController(userService *services.UserService, config config.Impl) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	services.UserService
	Config config.Impl
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/user", controller.Create)
	app.Post("/user/auth", controller.Authentication)
	app.Get("/user", middleware.AuthenticateUser(controller.Config.Get()), controller.Get)
}

func (controller *UserController) Get(c *fiber.Ctx) error {
	email := c.Locals("email").(string)
	user := controller.UserService.GetByEmail(email)
	return c.Status(200).JSON(user.ToResponse())
}

func (controller *UserController) Create(c *fiber.Ctx) error {
	var request models.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)
	if len(request.Roles) != 0 {
		if controller.Config.Get().Server.UserKey != request.UserKey {
			exception.PanicLogging("User key is not valid!")
		}
	}
	controller.UserService.Create(request.Username, request.Password, request.Email, request.Roles)
	return c.Status(200).JSON(fiber.Map{"message": "User created successfully!"})
}

func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request models.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Context(), request)
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Role,
		})
	}
	tokenJwtResult := common.GenerateToken(result.Email, userRoles, controller.Config.Get())
	resultWithToken := map[string]interface{}{
		"token":    tokenJwtResult,
		"username": result.Email,
		"role":     userRoles,
	}
	return c.Status(fiber.StatusOK).JSON(models.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}
