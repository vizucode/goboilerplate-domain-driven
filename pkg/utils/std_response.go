package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type customError struct {
	Msg  string
	Code int
}

func HandleError(msg string, code int) error {
	return &customError{
		Msg:  msg,
		Code: code,
	}
}

func (err *customError) Error() string {
	return err.Msg
}

func WriteError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	msg := "Internal Server Error"
	errorData := map[string]any{}

	if he, ok := err.(*customError); ok {
		status = he.Code
		msg = he.Msg
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		status = 400
		msg = "Please check your input"

		errorData = make(map[string]any)
		for _, val := range validationErrors {
			field := val.Field()
			tag := strings.ToLower(val.ActualTag())

			var msg string

			switch tag {
			case "required":
				msg = fmt.Sprintf("%s is required", field)

			case "min":
				msg = fmt.Sprintf("%s must be at least %s", field, val.Param())

			case "max":
				msg = fmt.Sprintf("%s must be at most %s", field, val.Param())

			case "len":
				msg = fmt.Sprintf("%s must be exactly %s characters", field, val.Param())

			case "email":
				msg = fmt.Sprintf("%s must be a valid email address", field)

			case "numeric":
				msg = fmt.Sprintf("%s must be numeric", field)

			case "oneof":
				msg = fmt.Sprintf("%s must be one of: %s", field, val.Param())

			case "gte":
				msg = fmt.Sprintf("%s must be greater than or equal to %s", field, val.Param())

			case "gt":
				msg = fmt.Sprintf("%s must be greater than %s", field, val.Param())

			case "lte":
				msg = fmt.Sprintf("%s must be less than or equal to %s", field, val.Param())

			case "lt":
				msg = fmt.Sprintf("%s must be less than %s", field, val.Param())

			default:
				msg = fmt.Sprintf("%s is not valid (%s)", field, tag)
			}

			errorData[strings.ToLower(field)] = msg
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "Error",
		"message": msg,
		"data":    nil,
		"error":   errorData,
	})
}

func WriteOK(w http.ResponseWriter, message string, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	statusMsg := "OK"
	if status == 201 {
		statusMsg = "Created"
	}

	json.NewEncoder(w).Encode(map[string]any{
		"status":  statusMsg,
		"message": message,
		"data":    data,
	})
}
