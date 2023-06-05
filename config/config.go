package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type DbConfig struct {
	DataSourceName string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() *Config {
	urlApi := os.Getenv("API_URL")
	apiPort := os.Getenv("API_PORT")

	c.ApiConfig = ApiConfig{
		ApiPort: apiPort,
		ApiHost: urlApi,
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbPort, dbSslMode)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
