package entity

import "github.com/google/uuid"

type Haendler struct {
	Id   uuid.UUID `gorm:"primaryKey;column:user_role_id;type:varchar(36)"`
	Name string    `gorm:"column:Name;type:varchar(100)"`
	Hsn  string    `gorm:"column:Hsn;type:varchar(100)"`
	Cars []Car     `gorm:"ForeignKey:Hsn;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Car struct {
	Id   uuid.UUID `gorm:"primaryKey;column:user_role_id;type:varchar(36)"`
	Name string    `gorm:"column:Name;type:varchar(100)"`
	Hsn  string    `gorm:"column:Hsn;type:varchar(100)"`
	Tsn  string    `gorm:"column:Tsn;type:varchar(100)"`
}
