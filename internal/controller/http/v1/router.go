package v1

import (
	"github.com/koizumi55555/corporation-api/config"
	"github.com/koizumi55555/corporation-api/internal/usecase"
	"github.com/koizumi55555/corporation-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *gin.Engine, cfg *config.Config,
	corporationUC usecase.CorporationUseCase,
	queueUC usecase.Queue, l *logger.Logger,
) error {
	v1h := handler.Group("/v1")
	NewCorporationRoutes(v1h, corporationUC, queueUC, l)
	return nil
}
