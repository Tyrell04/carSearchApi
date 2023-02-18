package config

import (
	"github.com/marcleonschulz/carSearchApi/exception"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Impl interface {
	Get() Config
}

type configImpl struct {
}

func New() Impl {
	return &configImpl{}
}

func (config *configImpl) Get() Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v := LoadConfig(cfgPath, "yml")
	return ParseConfig(v)
}

func ParseConfig(v *viper.Viper) Config {
	var cfg Config
	err := v.Unmarshal(&cfg)
	exception.PanicLogging(err)
	return cfg
}
func LoadConfig(filename string, fileType string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			exception.PanicLogging(err)
		}
	}
	return v
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "config/config-docker"
	} else {
		return "config/config-development"
	}
}
