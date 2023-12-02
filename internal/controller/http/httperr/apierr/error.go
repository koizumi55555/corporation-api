package apierr

type ApiErrF interface {
	Error() *ErrResponses
	StatusCode() int
}

type ErrorCodeInvalidRequest struct{}

func (e ErrorCodeInvalidRequest) StatusCode() int {
	return getStatusCode(e.Error().ErrorMessage)
}

func (e ErrorCodeInvalidRequest) Error() *ErrResponses {
	return &ErrResponses{
		ErrorMessage: ErrCodeInvalidRequest,
	}
}

type ErrorCodeValidationFailed struct{}

func (e ErrorCodeValidationFailed) StatusCode() int {
	return getStatusCode(e.Error().ErrorMessage)
}

func (e ErrorCodeValidationFailed) Error() *ErrResponses {
	return &ErrResponses{
		ErrorMessage: ErrCodeValidationFailed,
	}
}

type ErrorCodeResourceNotFound struct{}

func (e ErrorCodeResourceNotFound) StatusCode() int {
	return getStatusCode(e.Error().ErrorMessage)
}

func (e ErrorCodeResourceNotFound) Error() *ErrResponses {
	return &ErrResponses{
		ErrorMessage: ErrCodeResourceNotFound,
	}
}

type ErrorCodeInternalServerError struct{}

func (e ErrorCodeInternalServerError) StatusCode() int {
	return getStatusCode(e.Error().ErrorMessage)
}
func (e ErrorCodeInternalServerError) Error() *ErrResponses {
	return &ErrResponses{
		ErrorMessage: ErrCodeInternalServerError,
	}
}

type ErrorCodeConflict struct{}

func (e ErrorCodeConflict) StatusCode() int {
	return getStatusCode(e.Error().ErrorMessage)
}
func (e ErrorCodeConflict) Error() *ErrResponses {
	return &ErrResponses{
		ErrorMessage: ErrCodeConflict,
	}
}
