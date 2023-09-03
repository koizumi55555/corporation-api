package apierr

import (
	"fmt"
	"net/http"
)

const (
	// ErrCodeInvalidRequest : [400] Request is invalid
	ErrCodeInvalidRequest = "invalid_request"
	// ErrCodeValidationFailed : [400] Validation Failed
	ErrCodeValidationFailed = "validation_failed"
	// ErrCodeResourceNotFound : [404] Resource not found
	ErrCodeResourceNotFound = "resource_not_found"
	// ErrCodeConflict : [409] Conflict
	ErrCodeConflict = "conflict"
	// ErrCodeInternalServerError : [500] Internal Server Error
	ErrCodeInternalServerError = "internal_server_error"
)

func StatusCode(errCode string) (int, error) {
	statusCode := getStatusCode(errCode)
	if statusCode == -1 {
		return -1, fmt.Errorf("invalid err code: %s", errCode)
	}
	return statusCode, nil
}

func getStatusCode(errCode string) int {
	switch errCode {
	case ErrCodeInvalidRequest, ErrCodeValidationFailed:
		return http.StatusBadRequest
	case ErrCodeResourceNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	case ErrCodeInternalServerError:
		return http.StatusInternalServerError
	}
	return -1
}
