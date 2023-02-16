package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)

	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else {
		return "config-development"
	}
}
