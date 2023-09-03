package master_repo

import (
	"koizumi55555/corporation-api/build/db"
	"testing"

	"gorm.io/gorm"
)

func makeMasterRepo(t *testing.T) *MasterRepository {
	t.Helper()

	// l, err := logger.New(config.LoggerConfig{Encoding: "console", Level: "debug")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	dbh, err := db.NewDBHandler("localhost", "5432", "corporation-api-user",
		"corporation-api", "corporation-api-pw", "disable", "", 1)
	if err != nil {
		dbh, err = db.NewDBHandler("postgres", "5432", "corporation-api-user",
			"corporation-api", "corporation-api-pw", "disable", "", 1)
		if err != nil {
			t.Fatal(err)
		}
	}
	return New(dbh, l)
}

func TruncateDB(t *testing.T, db *db.DBHandler) {
	t.Helper()

	tables := []interface{}{}
	for _, table := range tables {
		if err := db.Conn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&table).Error; err != nil {
			t.Fatal(err)
		}
	}
}
