package services

import "github.com/marcleonschulz/carSearchApi/entity"

type CarService interface {
	GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler, error)
	GetByHsn(hsn string) (entity.Haendler, error)
	Create(hsn string, tsn string, name string, haendlerName string)
	CreateCarBulk(cars []entity.CarCreateBulk) error
}
