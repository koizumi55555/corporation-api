package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	Conn *gorm.DB
}

func NewDBHandler(host, port, user, password, db, sslMode string) (*DBHandler, error) {
	conn, err := connect(host, port, user, password, db, sslMode)
	if err != nil {
		return nil, err
	}
	dbh := &DBHandler{
		Conn: conn,
	}

	return dbh, nil
}

func connect(host, port, user, password, db, sslMode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, port, user, password, db, sslMode)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	return conn, nil
}
