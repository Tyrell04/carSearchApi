package main

import "github.com/marcleonschulz/carSearchApi/config"

func main() {
	cfg := config.GetConfig()
	println(cfg)
}
