package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading env files")
	}
}

func Env(key string) string {
	return os.Getenv(key)
}
