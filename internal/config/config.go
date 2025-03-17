package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnStr    string
	JWTSecretKey string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	return Config{
		DBConnStr:    os.Getenv("DB_CONN_STR"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
