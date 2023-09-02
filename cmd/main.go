package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/controller"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository"
	"github.com/marcleonschulz/carSearchApi/internal/repository/impl"
	services "github.com/marcleonschulz/carSearchApi/services/impl"
	"log"
	"os"
)

func main() {
	cfg := config.New().Get()
	repository.InitDb(&cfg)
	logFile, err := os.OpenFile("log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Logging to custom file")

	// register repositories
	userRepository := impl.NewUserRepositoryImpl(repository.GetDb())
	carRepository := impl.NewCarRepositoryImpl(repository.GetDb())

	// register services
	userService := services.NewUserServiceImpl(&userRepository)
	carService := services.NewCarServiceImpl(&carRepository)

	// register controllers
	userController := controller.NewUserController(&userService, config.New())
	carContoller := controller.NewCarController(&carService, config.New())
	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
			Prefork:      false,
		})
	app.Use(recover.New())
	app.Use(cors.New())

	userController.Route(app)
	carContoller.Route(app)
	log.Println("Run server on port: " + cfg.Server.Port + " ...")
	err = app.Listen(":" + cfg.Server.Port)
	exception.PanicLogging(err)
}
