package v1

import (
	"koizumi55555/corporation-api/internal/controller/http/httperr"
	"koizumi55555/corporation-api/internal/controller/http/model"
	"koizumi55555/corporation-api/internal/entity"
	"koizumi55555/corporation-api/internal/usecase"
	"koizumi55555/corporation-api/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type corporationRoutes struct {
	l             *logger.Logger
	corporationUC usecase.CorporationUseCase
}

func NewCorporationRoutes(handler *gin.RouterGroup,
	corporationUC usecase.CorporationUseCase, l *logger.Logger) {
	r := &corporationRoutes{l, corporationUC}
	handler.GET("/corporation/:corporation_id", r.GetCorporation)
	handler.GET("/corporation", r.GetCorporationList)
	handler.POST("/corporation", r.CreateCorporation)
	handler.PATCH("/corporation/:corporation_id", r.UpdateCorporation)
	handler.DELETE("/corporation/:corporation_id", r.DeleteCorporation)
}

// Get Corporation
func (r *corporationRoutes) GetCorporation(c *gin.Context) {
	// validation
	corpID, validationErr := model.ValidateCorporationIdRequest(c)
	if validationErr != nil {
		r.l.Warn(validationErr.Error().ErrorCode)
		httperr.ErrorResponse(c, validationErr)
		return
	}

	// Get Corporation
	corp, err := r.corporationUC.GetCorporation(c, corpID)
	if err != nil {
		r.l.Warn(err.Error().ErrorCode)
		httperr.ErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, makeCorporationResponse(corp))
}

// Get Corporation List
func (r *corporationRoutes) GetCorporationList(c *gin.Context) {
	// Get Corporation List
	corpList, err := r.corporationUC.GetCorporationList(c)
	if err != nil {
		r.l.Warn(err.Error().ErrorCode)
		httperr.ErrorResponse(c, err)
		return
	}

	// response
	c.JSON(http.StatusOK, makeCorporationResponse(corpList))
}

// Create Corporation
func (r *corporationRoutes) CreateCorporation(c *gin.Context) {
	// validation
	corporationPost, validationErr := model.ValidatePostCorporationRequest(c)
	if validationErr != nil {
		r.l.Warn(validationErr.Error().ErrorCode)
		httperr.ErrorResponse(c, validationErr)
		return
	}

	// entity作成
	input := entity.Corporation{
		CorporationID: "",
		Name:          corporationPost.Name,
		Domain:        corporationPost.Domain,
		Number:        corporationPost.Number,
		CorpType:      corporationPost.CorpType,
	}

	// Create Corporation
	corp, err := r.corporationUC.CreateCorporation(c, input)
	if err != nil {
		r.l.Warn(err.Error().ErrorCode)
		httperr.ErrorResponse(c, err)
		return
	}

	// response
	c.JSON(http.StatusCreated, makeCorporationResponse(corp))
}

// Update Corporation
func (r *corporationRoutes) UpdateCorporation(c *gin.Context) {
	// validation
	corpID, corporationPatch, validationErr := model.ValidatePatchCorporationRequest(c)
	if validationErr != nil {
		r.l.Warn(validationErr.Error().ErrorCode)
		httperr.ErrorResponse(c, validationErr)
		return
	}

	// entity作成
	input := entity.Corporation{
		CorporationID: corpID,
		Name:          *corporationPatch.Name,
		Domain:        *corporationPatch.Domain,
		Number:        *corporationPatch.Number,
		CorpType:      *corporationPatch.CorpType,
	}

	// Update Corporation
	corp, err := r.corporationUC.UpdateCorporation(c, input)
	if err != nil {
		r.l.Warn(err.Error().ErrorCode)
		httperr.ErrorResponse(c, err)
		return
	}

	// response
	c.JSON(http.StatusOK, makeCorporationResponse(corp))
}

// Delete Corporation
func (r *corporationRoutes) DeleteCorporation(c *gin.Context) {
	// validation
	corpID, validationErr := model.ValidateCorporationIdRequest(c)
	if validationErr != nil {
		r.l.Warn(validationErr.Error().ErrorCode)
		httperr.ErrorResponse(c, validationErr)
		return
	}
	err := r.corporationUC.DeleteCorporation(c, corpID)
	if err != nil {
		r.l.Warn(err.Error().ErrorCode)
		httperr.ErrorResponse(c, err)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func makeCorporationResponse(corp []entity.Corporation) []model.Corporation {
	corporations := make([]model.Corporation, len(corp))
	for i, c := range corp {
		corporations[i] = model.Corporation{
			CorporationId: c.CorporationID,
			Name:          c.Name,
			Domain:        c.Domain,
			Number:        c.Number,
			CorpType:      c.CorpType,
		}
	}
	return corporations
}
