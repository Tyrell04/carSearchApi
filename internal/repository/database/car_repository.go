package database

import "github.com/marcleonschulz/carSearchApi/entity"

type CarRepository interface {
	GetByHsnTsn(hsn string, tsn string) (entity.Car, entity.Haendler)
	Create(hsn string, tsn string, name string, haendlerName string)
}
