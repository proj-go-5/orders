package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	dsn string
}

func (db *Database) GetConnection() (*gorm.DB, func(), error) {
	connection, err := gorm.Open(postgres.Open(db.dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, nil, err
	}

	stop := func() {
		err := sqlDB.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
		fmt.Println("Database connection closed")
	}

	return connection, stop, nil
}

func NewDatabase(dsn string) *Database {
	return &Database{dsn}
}
