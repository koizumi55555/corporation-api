package master_repo

import (
	db "koizumi55555/corporation-api/build/db/sql"
	"testing"

	"gorm.io/gorm"
)

func makeMasterRepo(t *testing.T) *MasterRepository {
	t.Helper()
	dbh, err := db.NewDBHandler()
	if err != nil {
		dbh, err = db.NewDBHandler()
		if err != nil {
			t.Fatal(err)
		}
	}
	return New(dbh)
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
