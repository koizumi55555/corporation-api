package master_repo

import (
	"testing"

	"github.com/koizumi55555/corporation-api/internal/usecase/master_repo/schema"
	db "github.com/koizumi55555/corporation-api/pkg/rdb"

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

	tables := []interface{}{
		schema.Corporation{},
	}
	for _, table := range tables {
		if err := db.Conn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&table).Error; err != nil {
			t.Fatal(err)
		}
	}
}

func SeedData(t *testing.T, db *db.DBHandler) {
	t.Helper()

	seedFuncs := []func(db *gorm.DB) error{
		SeedCorporation,
	}

	for _, seedFunc := range seedFuncs {
		if err := seedFunc(db.Conn); err != nil {
			t.Fatal(err)
		}
	}
}
