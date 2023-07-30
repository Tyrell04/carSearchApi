package controller

import (
	"encoding/csv"
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/entity"
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
	app.Post("/car/bulk", middleware.AuthenticateRoles("admin", carController.Config.Get()), carController.CreateCarBulk)
}

func (carController *CarController) GetByHsnTsn(c *fiber.Ctx) error {
	hsn := c.Params("hsn")
	tsn := c.Params("tsn")
	car, haendler, err := carController.CarService.GetByHsnTsn(hsn, tsn)
	if err != nil {
		panic(exception.NotFoundError{Message: err.Error()})
	}
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
	haendler, err := carController.CarService.GetByHsn(hsn)
	if err != nil {
		panic(exception.NotFoundError{Message: err.Error()})
	}
	return c.JSON(fiber.Map{
		"haendler": haendler.ToResponse(),
	})
}

func (carController *CarController) CreateCarBulk(c *fiber.Ctx) error {
	cars := []entity.CarCreateBulk{}
	file, err := c.FormFile("file")
	exception.PanicLogging(err)
	f, err := file.Open()
	exception.PanicLogging(err)
	defer f.Close()
	exception.PanicLogging(err)
	lines, err := csv.NewReader(f).ReadAll()
	exception.PanicLogging(err)

	for _, line := range lines {
		cars = append(cars, entity.CarCreateBulk{
			Hsn:      line[0],
			Haendler: line[1],
			Tsn:      line[2],
			Name:     line[3],
		})
	}
	carController.CarService.CreateCarBulk(cars)
	return c.JSON(fiber.Map{"message": "Car created successfully!"})
}
