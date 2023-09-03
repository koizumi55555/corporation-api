package master_repo

import (
	db "koizumi55555/corporation-api/build/db/sql"
	"koizumi55555/corporation-api/pkg/logger"
)

type MasterRepository struct {
	DBHandler *db.DBHandler
	l         *logger.Logger
}

func New(dbh *db.DBHandler) *MasterRepository {
	return &MasterRepository{DBHandler: dbh}
}
