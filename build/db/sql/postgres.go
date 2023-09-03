package db

import (
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
	db := new(DBHandler)
	db.Conn = conn
	return db, nil
}

func connect() (*gorm.DB, error) {
	dsn := "host=localhost user=user1 password=password dbname=koizumi port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return conn, err
}
