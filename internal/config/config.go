package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Env    string
	DBUrl  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	return &Config{
		Port:  getEnv("APP_PORT", "8080"),
		Env:   getEnv("APP_ENV", "development"),
		DBUrl: getEnv("DB_URL", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal
	}
	return val
}