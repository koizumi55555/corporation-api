package v1

import (
	"koizumi55555/go-restapi/config"
	"koizumi55555/go-restapi/internal/usecase"
	"koizumi55555/go-restapi/pkg/logger"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *gin.Engine, cfg *config.Config,
	corporationUC usecase.CorporationUseCase, l *logger.Logger,
) error {
	v1h := handler.Group("/v1")
	NewCorporationRoutes(v1h, corporationUC, l)
	return nil
}
