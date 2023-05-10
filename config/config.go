package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	MongoDBHost     string
	MongoDBPort     string
	MongoDBDatabase string
	StoragePath     string
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Config{
		Port:            os.Getenv("PORT"),
		MongoDBHost:     os.Getenv("MONGODB_HOST"),
		MongoDBPort:     os.Getenv("MONGODB_PORT"),
		MongoDBDatabase: os.Getenv("MONGODB_DATABASE"),
		StoragePath:     os.Getenv("STORAGE_PATH"),
	}

	return config
}
