package services

import "github.com/marcleonschulz/carSearchApi/entity"

type CarService interface {
	GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler)
	Create(hsn string, tsn string, name string, haendlerName string)
}
