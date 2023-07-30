package entity

import (
	"github.com/marcleonschulz/carSearchApi/pkg/models"
)

type Haendler struct {
	Name string `gorm:"column:name;type:varchar(150)"`
	Hsn  string `gorm:"primaryKey;column:hsn;type:varchar(10)"`
	Cars []Car  `gorm:"ForeignKey:hsn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Car struct {
	Id   uint   `gorm:"primaryKey;column:id;type:int(11)"`
	Tsn  string `gorm:"column:tsn;type:varchar(10)"`
	Name string `gorm:"column:name;type:varchar(150)"`
	Hsn  string `gorm:"column:hsn;type:varchar(10)"`
}

type CarCreateBulk struct {
	Tsn      string
	Hsn      string
	Name     string
	Haendler string
}

func (Haendler) TableName() string {
	return "haendler"
}

func (Car) TableName() string {
	return "car"
}

func (car Car) ToResponse() models.CarResponse {
	return models.CarResponse{
		Tsn:  car.Tsn,
		Name: car.Name,
	}
}

func (haendler Haendler) ToResponse() models.HaendlerResponse {
	return models.HaendlerResponse{
		Name: haendler.Name,
		Hsn:  haendler.Hsn,
	}
}
