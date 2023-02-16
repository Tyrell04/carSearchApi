package database

import (
	"fmt"
	"github.com/marcleonschulz/carSearchApi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Berlin",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime * time.Minute)

	log.Println("Db connection established")
	err = dbClient.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	err := con.Close()
	if err != nil {
		log.Fatal("Error while shutdown: ", err)
	}
}
