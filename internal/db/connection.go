package db

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	dsn string
}

func (db *Database) GetConnection(ctx context.Context) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(db.dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		err := sqlDB.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	return connection, nil
}

func NewDatabase(dsn string) *Database {
	return &Database{dsn}
}
