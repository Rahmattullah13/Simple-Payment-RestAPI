package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     string
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type DbConfig struct {
	DataSourceName string
}
type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

func (c *Config) readConfig() *Config {
	urlApi := os.Getenv("API_URL")
	apiPort := os.Getenv("API_PORT")

	c.ApiConfig = ApiConfig{
		ApiPort: apiPort,
		ApiHost: urlApi,
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:  "ENIGMA",
		JwtSignatureKey:  "password",
		JwtSigningMethod: jwt.SigningMethodHS256,
		// temporary change access token time from 1min to 60min
		AccessTokenLifeTime: 60 * time.Minute,
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
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
