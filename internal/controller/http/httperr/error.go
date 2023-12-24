package httperr

import (
	"github.com/koizumi55555/corporation-api/internal/controller/http/httperr/apierr"

	"github.com/gin-gonic/gin"
)

type response struct {
	Error string `json:"error" example:"message"`
}

func ErrorResponse(c *gin.Context, e apierr.ApiErrF) {
	c.AbortWithStatusJSON(
		e.StatusCode(),
		e.Error(),
	)
}
