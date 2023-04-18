package entity

import "github.com/google/uuid"

type Haendler struct {
	Id   uuid.UUID `gorm:"primaryKey;column:haendler_id;type:varchar(36)"`
	Name string    `gorm:"column:name;type:varchar(100)"`
	Hsn  string    `gorm:"column:hsn;type:varchar(10)"`
	Cars []Car     `gorm:"ForeignKey:HaendlerId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Car struct {
	Id         uuid.UUID `gorm:"primaryKey;column:car_id;type:varchar(36)"`
	Tsn        string    `gorm:"column:tsn;type:varchar(10)"`
	Name       string    `gorm:"column:name;type:varchar(100)"`
	HaendlerId uuid.UUID `gorm:"column:haendler_id;type:varchar(36)"`
}
