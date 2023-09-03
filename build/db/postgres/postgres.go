package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func NewDBHandler() (*DBHandler, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	db := &DBHandler{
		Conn: conn,
	}
	return db, nil
}

func connect() (*gorm.DB, error) {
	dsn := "host=localhost user=corporation-api-user password=corporation-api-pw dbname=corporation-api port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	return conn, nil
}
