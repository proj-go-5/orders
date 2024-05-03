package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"orders/internal/config"
)

func GetConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env("DB_HOST"),
		config.Env("DB_USER"),
		config.Env("DB_PASSWORD"),
		config.Env("DB_NAME"),
	)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return connection
}
