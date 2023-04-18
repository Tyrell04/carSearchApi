package impl

import (
	"context"
	"errors"
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

func (userRepository *userRepositoryImpl) Create(username string, password string, email string, roles []string) {
	var userRoles []entity.UserRole
	for _, role := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Id:    uuid.New(),
			Email: email,
			Role:  role,
		})
	}
	user := entity.User{
		Username:  username,
		Password:  utils.HashPassword(password),
		Email:     email,
		IsActive:  true,
		UserRoles: userRoles,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) GetByEmail(email string) entity.User {
	var user entity.User
	err := userRepository.DB.Where("email = ?", email).First(&user).Error
	exception.PanicLogging(err)
	err = userRepository.DB.Model(&user).Association("UserRoles").Find(&user.UserRoles)
	exception.PanicLogging(err)
	return user
}

func (userRepository *userRepositoryImpl) GetByUsername(username string) entity.User {
	var user entity.User
	err := userRepository.DB.Where("username = ?", username).First(&user).Error
	exception.PanicLogging(err)
	err = userRepository.DB.Model(&user).Association("UserRoles").Find(&user.UserRoles)
	exception.PanicLogging(err)
	return user
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).
		Preload("UserRoles").
		Where("tb_user.email = ? and tb_user.is_active = ?", email, true).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
