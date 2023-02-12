package user

import (
	"github.com/google/uuid"
	"github.com/marcleonschulz/carSearchApi/internal/hash"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
	"gorm.io/gorm"
)

type UserService struct {
	database *gorm.DB
}

func NewUserService() *UserService {
	database := database.GetDb()
	return &UserService{
		database: database,
	}
}

func (s *UserService) CreateUser(user *database.User) (*database.User, error) {
	var err error
	if s.checkExistenceOfUser(user) {
		return nil, nil
	}
	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	err = s.database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserById(id uuid.UUID) (*database.User, error) {
	user := &database.User{}
	err := s.database.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserService) checkExistenceOfUser(user *database.User) bool {
	var exists bool
	if err := s.database.Model(&database.User{}).
		Select("count(*) > 0").
		Where("user = ?", user.Email).
		Find(&exists).
		Error; err != nil {
		return false
	}
	return exists
}
