package config

import "time"

type Config struct {
	Database Database    `mapstructure:"database"`
	Redis    RedisConfig `mapstructure:"redis"`
	Jwt      Jwt         `mapstructure:"jwt"`
	Server   Server      `mapstructure:"server"`
	Cors     CorsConfig  `mapstructure:"cors"`
}

type Server struct {
	Port string `mapstructure:"port"`
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

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type Jwt struct {
	Secret string
}
