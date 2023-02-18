package database

import "github.com/marcleonschulz/carSearchApi/entity"

type UserRepository interface {
	GetByEmail(email string) entity.User
	Create(username string, password string, roles []string)
}
