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
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	Password           string        `mapstructure:"password"`
	Db                 string        `mapstructure:"db"`
	DialTimeout        time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout        time.Duration `mapstructure:"read_timeout"`
	WriteTimeout       time.Duration `mapstructure:"write_timeout"`
	IdleCheckFrequency time.Duration `mapstructure:"idle_check_frequency"`
	PoolSize           int           `mapstructure:"pool_size"`
	PoolTimeout        time.Duration `mapstructure:"pool_timeout"`
}

type Jwt struct {
	Secret                    string        `mapstructure:"secret"`
	AccessTokenExpireDuration time.Duration `mapstructure:"access_token_expire_duration"`
}
