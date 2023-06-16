package impl

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
	"gorm.io/gorm"
)

func NewCarRepositoryImpl(DB *gorm.DB) database.CarRepository {
	return &carRepositoryImpl{DB: DB}
}

type carRepositoryImpl struct {
	*gorm.DB
}

func (carRepository *carRepositoryImpl) GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler, error) {
	var car entity.Car
	var haendler entity.Haendler
	result := carRepository.DB.Where("hsn = ?", hsn).First(&haendler)
	fmt.Println(hsn)
	result = carRepository.DB.Where("tsn = ? AND haendler_id = ?", tsn, haendler.Id).First(&car)
	if result.RowsAffected == 0 {
		return entity.Car{}, entity.Haendler{}, errors.New("Not Found")
	}
	return car, haendler, nil
}

func (carRepository *carRepositoryImpl) GetByHsn(hsn string) (entity.Haendler, error) {
	var haendler entity.Haendler
	result := carRepository.DB.Where("hsn = ?", hsn).Find(&haendler)
	if result.RowsAffected == 0 {
		return entity.Haendler{}, errors.New("Not Found")
	}
	return haendler, nil
}

func (carRepository *carRepositoryImpl) Create(hsn string, tsn string, name string, haendlerName string) {
	var car entity.Car
	var haendler entity.Haendler

	result := carRepository.DB.Where("name = ?", haendlerName).First(&haendler)
	if result.RowsAffected == 0 {
		haendler = entity.Haendler{
			Id:   uuid.New(),
			Name: haendlerName,
			Hsn:  hsn,
		}
		err := carRepository.DB.Create(&haendler).Error
		exception.PanicLogging(err)
	}
	err := carRepository.DB.Where("haendler_id = ? AND tsn = ?", haendler.Id, tsn).First(&car).Error
	if err != nil {
		car = entity.Car{
			Id:         uuid.New(),
			HaendlerId: haendler.Id,
			Tsn:        tsn,
			Name:       name,
		}
		//create car with the association to the haendler
		err = carRepository.DB.Model(&haendler).Association("Cars").Append(&car)
		exception.PanicLogging(err)
	} else {
		exception.PanicLogging(errors.New("Car already exists"))
	}
}
