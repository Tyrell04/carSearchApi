package services

import (
	"context"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
)

type UserService interface {
	GetByEmail(email string) entity.User
	GetByUsername(username string) entity.User
	Create(username string, password string, email string, roles []string)
	Authentication(ctx context.Context, model models.UserModel) entity.User
}
