package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI        string
	DBName       string
	MongoTimeOut time.Duration
	ServerPort   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}

	dbURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")
	mongoTimeout, _ := time.ParseDuration(os.Getenv("MONGODB_TIMEOUT"))
	serverPort := os.Getenv("SERVER_PORT")

	return &Config{
		DBURI:        dbURI,
		DBName:       dbName,
		MongoTimeOut: mongoTimeout,
		ServerPort:   serverPort,
	}, nil

}
