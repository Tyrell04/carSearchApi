package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/controller"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository/cache"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
	repository "github.com/marcleonschulz/carSearchApi/internal/repository/database/impl"
	services "github.com/marcleonschulz/carSearchApi/services/impl"
)

func main() {
	cfg := config.New().Get()
	database.InitDb(&cfg)
	cache.InitRedis(&cfg)

	// register repositories
	userRepository := repository.NewUserRepositoryImpl(database.GetDb())

	// register services
	userService := services.NewUserServiceImpl(&userRepository)

	// register controllers
	userController := controller.NewUserController(&userService, config.New())

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	userController.Route(app)
	err := app.Listen(":" + cfg.Server.Port)
	exception.PanicLogging(err)
}
