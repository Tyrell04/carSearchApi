package impl

import (
	"errors"
	"fmt"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository"
	"github.com/marcleonschulz/carSearchApi/pkg/helper"
	"gorm.io/gorm"
	"strings"
)

func NewCarRepositoryImpl(DB *gorm.DB) repository.CarRepository {
	return &carRepositoryImpl{DB: DB}
}

type carRepositoryImpl struct {
	*gorm.DB
}

func (carRepository *carRepositoryImpl) GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler, error) {
	var car entity.Car
	var haendler entity.Haendler
	hsn = strings.ToUpper(hsn)
	tsn = strings.ToUpper(tsn)
	result := carRepository.DB.Where("hsn = ?", hsn).First(&haendler)
	fmt.Println(hsn)
	result = carRepository.DB.Where("tsn = ? AND hsn = ?", tsn, haendler.Hsn).First(&car)
	if result.RowsAffected == 0 {
		return entity.Car{}, entity.Haendler{}, errors.New("Not Found")
	}
	return car, haendler, nil
}

func (carRepository *carRepositoryImpl) GetByHsn(hsn string) (entity.Haendler, error) {
	var haendler entity.Haendler
	hsn = strings.ToUpper(hsn)
	result := carRepository.DB.Where("hsn = ?", hsn).Find(&haendler)
	if result.RowsAffected == 0 {
		return entity.Haendler{}, errors.New("Not Found")
	}
	return haendler, nil
}

func (carRepository *carRepositoryImpl) CreateCarBulk(cars []entity.CarCreateBulk) error {
	var dbHaendler []entity.Haendler
	var haendlerHsn []string
	err := carRepository.DB.Model(&entity.Haendler{}).Pluck("hsn", &haendlerHsn).Error
	exception.PanicLogging(err)
	for _, car := range cars {
		if !helper.BinaryFindString(haendlerHsn, car.Hsn) {
			dbHaendler = append(dbHaendler, entity.Haendler{
				Name: car.Haendler,
				Hsn:  car.Hsn,
			})
			haendlerHsn = append(haendlerHsn, car.Hsn)
		}
	}
	err = carRepository.DB.Create(&dbHaendler).Error
	exception.PanicLogging(err)
	var result []entity.Car
	err = carRepository.DB.Table("car").Select("tsn", "hsn").Find(&result).Error
	exception.PanicLogging(err)
	var createCars []entity.Car
	for _, car := range cars {
		isCar := false
		for _, dbCars := range result {
			if car.Hsn == dbCars.Hsn && car.Tsn == dbCars.Tsn {
				isCar = true
				break
			}
		}
		if isCar == false {
			createCars = append(createCars, entity.Car{
				Hsn:  car.Hsn,
				Tsn:  car.Tsn,
				Name: car.Name,
			})
		}
	}
	err = carRepository.DB.Create(&createCars).Error
	exception.PanicLogging(err)
	return nil
}

func (carRepository *carRepositoryImpl) Create(hsn string, tsn string, name string, haendlerName string) {
	var car entity.Car
	var haendler entity.Haendler

	result := carRepository.DB.Where("name = ?", haendlerName).First(&haendler)
	if result.RowsAffected == 0 {
		haendler = entity.Haendler{
			Name: haendlerName,
			Hsn:  hsn,
		}
		err := carRepository.DB.Create(&haendler).Error
		exception.PanicLogging(err)
	}
	err := carRepository.DB.Where("hsn = ? AND tsn = ?", haendler.Hsn, tsn).First(&car).Error
	if err != nil {
		car = entity.Car{
			Hsn:  hsn,
			Tsn:  tsn,
			Name: name,
		}
		//create car with the association to the haendler
		err = carRepository.DB.Model(&haendler).Association("Cars").Append(&car)
		exception.PanicLogging(err)
	} else {
		exception.PanicLogging(errors.New("car already exists"))
	}
}

func (carRepository *carRepositoryImpl) GetCarsByHsnTsn(condition map[string]interface{}) []entity.Car {
	var cars []entity.Car
	err := carRepository.DB.Where(condition).Find(&cars).Error
	exception.PanicLogging(err)
	return cars
}

func (carRepository *carRepositoryImpl) GetCarsArrayHsn(hsn string, cars []entity.Car) []entity.Car {
	var result []entity.Car
	for _, car := range cars {
		if car.Hsn == hsn {
			result = append(result, car)
		}
	}
	return result
}
