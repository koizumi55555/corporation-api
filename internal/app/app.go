package app

import (
	"fmt"
	"koizumi55555/go-restapi/build/db"
	"koizumi55555/go-restapi/config"
	v1 "koizumi55555/go-restapi/internal/controller/http/v1"
	"koizumi55555/go-restapi/internal/usecase"
	master_repo "koizumi55555/go-restapi/internal/usecase/master_repo"
	"koizumi55555/go-restapi/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {
	l := logger.New(cfg.Level)

	masterDBH, err := db.NewDBHandler()
	if err != nil {
		return fmt.Errorf("DBHandler error: %w", err)
	}

	mRepo := master_repo.New(masterDBH)
	corporationUC := usecase.NewCorporationUsecase(mRepo)
	handler := gin.New()
	if err := v1.NewRouter(handler, cfg, corporationUC, l); err != nil {
		return fmt.Errorf("/v1 handler error: %w", err)
	}
	return nil
}
