package services

import "github.com/marcleonschulz/carSearchApi/entity"

type UserService interface {
	GetByEmail(email string) entity.User
	Create(username string, password string, roles []string)
}
