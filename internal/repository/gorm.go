package repository

import (
	"fmt"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/entity"
	"github.com/marcleonschulz/carSearchApi/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) {
	var err error
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Berlin",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{
		CreateBatchSize: 1000,
	})

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()

	sqlDb.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime * time.Minute)

	log.Println("Db connection established")
	err = dbClient.AutoMigrate(&entity.User{}, &entity.UserRole{}, &entity.Haendler{}, &entity.Car{})
	exception.PanicLogging(err)

}

func GetDb() *gorm.DB {
	return dbClient
}
