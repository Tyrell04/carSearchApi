package impl

import (
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/internal/repository"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewCarServiceImpl(carRepository *repository.CarRepository) services.CarService {
	return &carServiceImpl{CarRepository: *carRepository}
}

type carServiceImpl struct {
	repository.CarRepository
}

func (carService *carServiceImpl) GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler, error) {
	return carService.CarRepository.GetByHsnTsn(hsn, tsn)
}

func (carService *carServiceImpl) Create(hsn string, tsn string, name string, haendlerName string) {
	carService.CarRepository.Create(hsn, tsn, name, haendlerName)
}

func (carService *carServiceImpl) GetByHsn(hsn string) (entity.Haendler, error) {
	return carService.CarRepository.GetByHsn(hsn)
}

func (carService *carServiceImpl) CreateCarBulk(cars []entity.CarCreateBulk) error {
	return carService.CarRepository.CreateCarBulk(cars)
}
