package apiErrors

import "net/http"

// getFirstOrNil helper per estrarre il primo elemento opzionale
func getFirstOrNil(data []any) any {
	if len(data) > 0 {
		return data[0]
	}
	return nil
}

// BadRequestError crea un errore 400 Bad Request
func BadRequestError(message string, errorCode ErrorCode, data ...any) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		ErrorCode:  errorCode,
		Message:    message,
		Data:       getFirstOrNil(data),
	}
}

// NewUnauthorizedError crea un errore 401 Unauthorized
func UnauthorizedError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusUnauthorized,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewForbiddenError crea un errore 403 Forbidden
func ForbiddenError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusForbidden,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewNotFoundError crea un errore 404 Not Found
func NotFoundError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusNotFound,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewConflictError crea un errore 409 Conflict
func ConflictError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusConflict,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewUnprocessableEntityError crea un errore 422 Unprocessable Entity
func UnprocessableEntityError(message string, errorCode ErrorCode, data ...any) *APIError {
	return &APIError{
		StatusCode: http.StatusUnprocessableEntity,
		ErrorCode:  errorCode,
		Message:    message,
		Data:       getFirstOrNil(data),
	}
}

// NewTooManyRequestsError crea un errore 429 Too Many Requests
func TooManyRequestsError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusTooManyRequests,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewInternalServerError crea un errore 500 Internal Server Error
func InternalServerError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusInternalServerError,
		ErrorCode:  errorCode,
		Message:    message,
	}
}

// NewServiceUnavailableError crea un errore 503 Service Unavailable
func ServiceUnavailableError(message string, errorCode ErrorCode) *APIError {
	return &APIError{
		StatusCode: http.StatusServiceUnavailable,
		ErrorCode:  errorCode,
		Message:    message,
	}
}
