package master_repo

import (
	"github.com/koizumi55555/corporation-api/pkg/logger"
	db "github.com/koizumi55555/corporation-api/pkg/rdb"
)

type MasterRepository struct {
	DBHandler *db.DBHandler
	l         *logger.Logger
}

func New(dbh *db.DBHandler) *MasterRepository {
	return &MasterRepository{DBHandler: dbh}
}
