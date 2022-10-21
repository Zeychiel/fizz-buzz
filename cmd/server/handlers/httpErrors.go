package handlers

import (
	"encoding/json"
	"net/http"
)

// A APIError is an error that is used when an API Handler needs to return an error.
// swagger:response APIError
type APIError struct {
	// The error message
	// Required: true
	// Example: Expected type int
	Message string `json:"message,omitempty"`
	// The code of the error
	// Required: true
	// Example: 404
	Code int
}

// Standard errors
var (
	// Code: 400
	// Message: Bad request
	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	// Code: 401
	// Message: Unauthorized
	ErrUnauthorized = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	// Code: 403
	// Message: Forbidden
	ErrForbidden = APIError{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}
	// Code: 404
	// Message: Not Found
	ErrNotFound = APIError{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
	// Code: 500
	// Message: Internal Server Error
	ErrInternalServerError = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	// Code: 502
	// Message: Bad Gateway
	ErrBadGateway = APIError{
		Code:    http.StatusBadGateway,
		Message: "Bad Gateway",
	}
	ErrServiceUnavailable = APIError{
		Code:    http.StatusServiceUnavailable,
		Message: "Service Unavailable",
	}
	ErrGatewayTimeout = APIError{
		Code:    http.StatusGatewayTimeout,
		Message: "Gateway Timeout",
	}
)

func (a *APIError) Throw(w http.ResponseWriter) {
	http.Error(w, a.Message, a.Code)
}

// NewAPIError returns a new *APIError
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Message: message,
		Code:    code,
	}
}

// ThrowCustom throws an APIError containing a custom message.
// This function is used to throw standard errors with an error used for debugging.
func (a *APIError) ThrowCustom(w http.ResponseWriter, message string) {
	http.Error(w, message, a.Code)
}

func WriteResponse(w http.ResponseWriter, response interface{}) {
	if response != nil {
		marshalled, err := json.Marshal(response)
		if err != nil {
			ErrInternalServerError.Throw(w)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(marshalled)
		return
	}
	w.WriteHeader(http.StatusOK)
}
