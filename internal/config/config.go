package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load(".env.local", ".env"); err != nil {
		fmt.Println("Error loading env files")
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
