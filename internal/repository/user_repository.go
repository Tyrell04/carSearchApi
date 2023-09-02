package repository

import (
	"context"
	"github.com/marcleonschulz/carSearchApi/entity"
)

type UserRepository interface {
	GetByEmail(email string) entity.User
	GetByUsername(username string) entity.User
	Create(username string, password string, email string, roles []string)
	Authentication(ctx context.Context, email string) (entity.User, error)
}
