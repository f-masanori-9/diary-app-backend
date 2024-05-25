package config

import (
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	APIPort *string
	DB      *DBConfig
}

func NewConfig() *Config {
	port := os.Getenv("API_PORT")
	return &Config{
		APIPort: &port,
	}
}
