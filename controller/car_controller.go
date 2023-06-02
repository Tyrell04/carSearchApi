package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/middleware"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewCarController(carService *services.CarService, config config.Impl) *CarController {
	return &CarController{CarService: *carService, Config: config}
}

type CarController struct {
	services.CarService
	Config config.Impl
}

func (carController *CarController) Route(app *fiber.App) {
	app.Get("/car/:hsn/:tsn", carController.GetByHsnTsn)
	app.Post("/car", middleware.AuthenticateRoles("admin", carController.Config.Get()), carController.Create)
}

func (carController *CarController) GetByHsnTsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	tsn := c.Params("tsn")
	car, haendler := carController.CarService.GetByHsnTsn(hsn, tsn)
	return c.JSON(fiber.Map{
		"car":      car.ToResponse(),
		"haendler": haendler.ToResponse(),
	})
}

func (carController *CarController) Create(c *fiber.Ctx) error {
	car := new(models.CarModel)
	err := c.BodyParser(car)
	exception.PanicLogging(err)
	carController.CarService.Create(car.Hsn, car.Tsn, car.Name, car.HaendlerName)
	return c.JSON(fiber.Map{"message": "Car created successfully!"})
}

func (carController *CarController) GetByHsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	haendler := carController.CarService.GetByHsn(hsn)
	return c.JSON(fiber.Map{
		"haendler": haendler.ToResponse(),
	})
}
