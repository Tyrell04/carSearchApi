package impl

import (
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
	"github.com/marcleonschulz/carSearchApi/services"
)

func NewUserServiceImpl(userRepository *database.UserRepository) services.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	database.UserRepository
}

func (userService *userServiceImpl) Create(username string, password string, roles []string) {
	userService.UserRepository.Create(username, password, roles)
}

func (userService *userServiceImpl) GetByEmail(email string) entity.User {
	return userService.UserRepository.GetByEmail(email)
}
