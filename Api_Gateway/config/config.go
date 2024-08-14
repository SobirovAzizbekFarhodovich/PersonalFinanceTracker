package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT     string
	BUDGETING_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	LOG_PATH string
	TokenKey string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(getEnv("AUTH_PORT", ":8002"))
	config.BUDGETING_PORT = cast.ToString(getEnv("BUDGETING_PORT", ":50055"))
	config.LOG_PATH = cast.ToString(getEnv("LOG_PATH", "logs/info.log"))
	config.TokenKey = cast.ToString(getEnv("TOKEN_KEY", "my_secret_key"))

	return config
}

func getEnv(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
