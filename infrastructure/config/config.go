package config

import (
	"log"
	"os"
)

type Configuration struct {
	AppConfig     AppConfig
	DbConfig      DataBaseConfig
	RedisDbConfig RedisDbConfig
}

type AppConfig struct {
	Port string
	Mode string
}

type DataBaseConfig struct {
	Host     string
	DbName   string
	User     string
	Password string
	Port     string
}

type RedisDbConfig struct {
	Address  string
	Password string
}

func NewAppConfig() *Configuration {
	config := &Configuration{
		AppConfig: AppConfig{
			Port: os.Getenv("APP_PORT"),
			Mode: os.Getenv("APP_MODE"),
		},
		DbConfig: DataBaseConfig{
			Host:     os.Getenv("DB_HOST"),
			DbName:   os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Port:     os.Getenv("DB_PORT"),
		},
		RedisDbConfig: RedisDbConfig{
			Address:  os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	log.Printf("DBNAME: %s", config.DbConfig.DbName)

	return config
}

func (a AppConfig) IsDebugMode() bool {
	return a.Mode != "prod" && a.Mode != "stg"
}
