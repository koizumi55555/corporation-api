package apierr

type ApiErrF interface {
	Error() *ErrResponses
	StatusCode() int
}

type ErrorCodeInvalidRequest struct{}

func (e ErrorCodeInvalidRequest) StatusCode() int {
	return getStatusCode(e.Error().ErrorCode)
}

func (e ErrorCodeInvalidRequest) Error() *ErrResponses {
	return &ErrResponses{
		ErrorCode: ErrCodeInvalidRequest,
	}
}

type ErrorCodeValidationFailed struct{}

func (e ErrorCodeValidationFailed) StatusCode() int {
	return getStatusCode(e.Error().ErrorCode)
}

func (e ErrorCodeValidationFailed) Error() *ErrResponses {
	return &ErrResponses{
		ErrorCode: ErrCodeValidationFailed,
	}
}

type ErrorCodeResourceNotFound struct{}

func (e ErrorCodeResourceNotFound) StatusCode() int {
	return getStatusCode(e.Error().ErrorCode)
}

func (e ErrorCodeResourceNotFound) Error() *ErrResponses {
	return &ErrResponses{
		ErrorCode: ErrCodeResourceNotFound,
	}
}

type ErrorCodeInternalServerError struct{}

func (e ErrorCodeInternalServerError) StatusCode() int {
	return getStatusCode(e.Error().ErrorCode)
}
func (e ErrorCodeInternalServerError) Error() *ErrResponses {
	return &ErrResponses{
		ErrorCode: ErrCodeInternalServerError,
	}
}

type ErrorCodeConflict struct{}

func (e ErrorCodeConflict) StatusCode() int {
	return getStatusCode(e.Error().ErrorCode)
}
func (e ErrorCodeConflict) Error() *ErrResponses {
	return &ErrResponses{
		ErrorCode: ErrCodeConflict,
	}
}
