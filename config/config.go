package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	DBName     string
	DBUser     string
	DBAddr     string
	DBHost     string
	DBPassword string
	DBPort     string
	REDISPort  string
	REDISHost  string
	JWTSecret  string
	Port       string
}

func InitConfig() *Config {
	loadAppEnv(os.Getenv("APP_ENV"))

	return &Config{
		AppEnv:     os.Getenv("APP_ENV"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBAddr:     os.Getenv("DB_ADDR"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
		REDISPort:  os.Getenv("REDIS_PORT"),
		REDISHost:  os.Getenv("REDIS_HOST"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		Port:       os.Getenv("PORT"),
	}
}

func loadAppEnv(env string) {
	if env == "" {
		env = "develop"
	}
	if env == "develop" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
