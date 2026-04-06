package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBIP           string
	DBPort         string
	DBName         string
	DBUser         string
	DBPassword     string
	ServerPort     string
	RedisAddr      string
	SendgridAPIKey string
}

func LoadConfig() *Config {

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("cannot load the env file")
	}

	return &Config{
		DBIP:           os.Getenv("DB_IP"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		SendgridAPIKey: os.Getenv("SENDGRID_API_KEY"),
	}
}
