package api

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents the structure of an error in the response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Stack   string `json:"stack"`
}

// Response represents the structure of a standard response
type Response struct {
	Success bool            `json:"success"`
	Errors  []ErrorResponse `json:"errors"`
	Data    interface{}     `json:"data"`
}

// WriteSuccessResponse writes a success response to the http.ResponseWriter
func WriteSuccessResponse(w http.ResponseWriter, data interface{}) {
	response := Response{
		Success: true,
		Errors:  []ErrorResponse{},
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// WriteFailureResponse writes a failure response to the http.ResponseWriter
func WriteFailureResponse(w http.ResponseWriter, statusCode int, errors ...ErrorResponse) {
	response := Response{
		Success: false,
		Errors:  errors,
		Data:    nil,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// NewErrorResponse creates a new ErrorResponse instance
func NewErrorResponse(message string, code string, stack string) ErrorResponse {
	return ErrorResponse{
		Message: message,
		Code:    code,
		Stack:   stack,
	}
}