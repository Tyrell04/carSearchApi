package impl

import (
	"github.com/google/uuid"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
	"github.com/marcleonschulz/carSearchApi/internal/utils"
	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) database.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Create(username string, password string, roles []string) {
	var userRoles []entity.UserRole
	for _, role := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Id:       uuid.New(),
			Username: username,
			Role:     role,
		})
	}
	user := entity.User{
		Username:  username,
		Password:  utils.HashPassword(password),
		IsActive:  true,
		UserRoles: userRoles,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) GetByEmail(email string) entity.User {
	var user entity.User
	err := userRepository.DB.Where("username = ?", email).First(&user).Error
	exception.PanicLogging(err)
	return user
}
