package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/internal/repository/cache"
	"github.com/marcleonschulz/carSearchApi/internal/server"
	"log"

	//"github.com/marcleonschulz/carSearchApi/internal/repository/cache"
	"github.com/marcleonschulz/carSearchApi/internal/repository/database"
)

func main() {
	if fiber.IsChild() {
		log.Println("Child process started")
	} else {
		log.Println("Main process started")
		cfg := config.GetConfig()
		err := cache.InitRedis(cfg)
		defer cache.CloseRedis()
		if err != nil {
			panic(err)
		}
		err = database.InitDb(cfg)
		defer database.CloseDb()
		if err != nil {
			panic(err)
		}
		server.NewServer(cfg)
	}

}
