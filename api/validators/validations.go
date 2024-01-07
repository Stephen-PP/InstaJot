package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Define our validator singleton
var validate = validator.New(validator.WithRequiredStructEnabled())

// Define ALL API types here
type (
	// User registration struct
	UserRegistration struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// Note Creation struct
	NoteCreation struct {
		NoteId  string `json:"note_id" validate:""`
		UserId  string `json:"user_id" validate:"required"`
		Content string `json:"content" validate:"required,hexadecimal"`
	}
)

// Define global API types here
type (
	EmptyStruct struct{}

	SuccessResponse struct {
		Success bool        `json:"success" validate:"eq=true"`
		Data    interface{} `json:"data"`
		Error   EmptyStruct `json:"error"`
	}

	FailureResponse struct {
		Success bool        `json:"success" validate:"eq=false"`
		Data    EmptyStruct `json:"data"`
		Error   string      `json:"error"`
	}
)

// Error response type (for Validate function)
type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func Validate(data interface{}) []ErrorResponse {
	// Create an array to store any validation errors
	validationErrors := []ErrorResponse{}

	// Validate
	errs := validate.Struct(data)
	if errs != nil {
		// Go through each validation error and create an ErrorResponse object
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			// Append the ErrorResponse object to the array
			validationErrors = append(validationErrors, elem)
		}
	}

	// Return the array of validation errors (if any)
	return validationErrors
}

func ParseValidationErrors(validationErrors []ErrorResponse) string {
	// Make sure the items in the array are Errors
	if len(validationErrors) == 0 || !validationErrors[0].Error {
		return "Unknown errors"
	}

	// Create an array to store all our error messages
	errMsgs := make([]string, 0)

	// Go through each validation error and create a string representing the error
	for _, err := range validationErrors {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	// Join all the error messages together and return them
	return strings.Join(errMsgs, " and ")
}
