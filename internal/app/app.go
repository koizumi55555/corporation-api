package app

import (
	"fmt"

	"github.com/koizumi55555/corporation-api/config"
	v1 "github.com/koizumi55555/corporation-api/internal/controller/http/v1"
	"github.com/koizumi55555/corporation-api/internal/usecase"
	master_repo "github.com/koizumi55555/corporation-api/internal/usecase/master_repo"
	"github.com/koizumi55555/corporation-api/internal/usecase/queue"
	"github.com/koizumi55555/corporation-api/pkg/logger"
	db "github.com/koizumi55555/corporation-api/pkg/rdb"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) error {
	l := logger.New(cfg.Level)

	masterDBH, err := db.NewDBHandler()
	if err != nil {
		return fmt.Errorf("DBHandler error: %w", err)
	}

	mRepo := master_repo.New(masterDBH)
	qu := queue.NewQueueUsecase(l, &cfg.QueueConfig)
	corporationUC := usecase.NewCorporationUsecase(mRepo)
	handler := gin.New()
	if err := v1.NewRouter(handler, cfg, corporationUC, qu, l); err != nil {
		return fmt.Errorf("/v1 handler error: %w", err)
	}
	err = handler.Run(":8080")
	if err != nil {
		return err
	}

	return nil
}
