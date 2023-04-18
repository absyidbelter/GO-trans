package config

import (
	"GO-Payment/pkg/utils"
	"time"
)

type DbConfig struct {
	Name           string
	Host, Port     string
	User, Password string
	SslMode        string
}

type ApiConfig struct {
	Host, Port string
}

type Storage struct {
	Dir string
}

type Secret struct {
	Key []byte
	Exp time.Duration
}

type AppConfig struct {
	ApiConfig ApiConfig
	DbConfig  DbConfig
	Secret    Secret
	Storage   Storage
}

func NewAppConfig() AppConfig {
	exp := time.Hour * 3

	return AppConfig{
		DbConfig: DbConfig{
			Name:     utils.GetEnv("DB_NAME"),
			Host:     utils.GetEnv("DB_HOST"),
			Port:     utils.GetEnv("DB_PORT"),
			User:     utils.GetEnv("DB_UNAME"),
			Password: utils.GetEnv("DB_PASSW"),
			SslMode:  utils.GetEnv("DB_SSL_MODE"),
		},
		ApiConfig: ApiConfig{
			Host: utils.GetEnv("HTTP_SERVER_HOST"),
			Port: utils.GetEnv("HTTP_SERVER_PORT"),
		},
		Secret: Secret{
			Key: []byte("SECRET_KEY"),
			Exp: time.Duration(exp.Seconds()) * time.Second,
		},
	}
}
