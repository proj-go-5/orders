package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
