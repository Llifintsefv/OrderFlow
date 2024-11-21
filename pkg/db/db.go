package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbConnection() (*gorm.DB, error) {
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	
	return db, nil
}