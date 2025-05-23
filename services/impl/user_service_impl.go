package impl

import (
	"context"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository"
	"github.com/marcleonschulz/carSearchApi/internal/utils"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) services.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Create(username string, password string, email string, roles []string) {
	userService.UserRepository.Create(username, password, email, roles)
}

func (userService *userServiceImpl) GetByEmail(email string) entity.User {
	return userService.UserRepository.GetByEmail(email)
}

func (userService *userServiceImpl) GetByUsername(username string) entity.User {
	return userService.UserRepository.GetByUsername(username)
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model models.UserModel) entity.User {
	userResult, err := userService.UserRepository.Authentication(ctx, model.Email)
	if !utils.CheckPasswordHash(model.Password, userResult.Password) || err != nil {
		panic(exception.UnauthorizedError{
			Message: "Invalid credentials",
		})
	}
	return userResult
}
