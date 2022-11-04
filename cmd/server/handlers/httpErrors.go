package handlers

import (
	"encoding/json"
	"log"
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

func WriteResponse(w http.ResponseWriter, response interface{}) {
	if response != nil {
		marshalled, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
			NewAPIError(http.StatusInternalServerError, err.Error()).Throw(w)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(marshalled); err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
