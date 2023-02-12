package database

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Role      bool      `gorm:"type:varchar(100);not null"`
}
