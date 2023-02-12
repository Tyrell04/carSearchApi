package database

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar(100);not null" json:"first_name""`
	LastName  string `gorm:"type:varchar(100);not null" json:"last_name"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string `gorm:"type:varchar(100);not null" json:"password"`
	Role      bool   `gorm:"type:bool;not null" json:"role"`
}
