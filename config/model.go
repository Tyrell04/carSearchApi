package config

import "time"

type Config struct {
	Database Database   `mapstructure:"database"`
	Jwt      Jwt        `mapstructure:"jwt"`
	Server   Server     `mapstructure:"server"`
	Cors     CorsConfig `mapstructure:"cors"`
}

type Server struct {
	Port        string `mapstructure:"port"`
	UserEnabled bool   `mapstructure:"user_enabled"`
	UserKey     string `mapstructure:"user_key"`
}

type CorsConfig struct {
	AllowOrigins string
}

type Database struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	Name            string        `mapstructure:"dbName"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type Jwt struct {
	Secret                    string        `mapstructure:"secret"`
	AccessTokenExpireDuration time.Duration `mapstructure:"access_token_expire_duration"`
}
