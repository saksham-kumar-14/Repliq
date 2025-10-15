package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	if intVal, err := strconv.Atoi(val); err == nil {
		return intVal
	}

	return fallback
}

func GetBool(key string, fallback bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return fallback
	}

	return boolVal
}
