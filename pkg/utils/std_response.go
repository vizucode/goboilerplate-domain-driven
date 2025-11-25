package utils

import (
	"encoding/json"
	"net/http"
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

	if he, ok := err.(*customError); ok {
		status = he.Code
		msg = he.Msg
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]any{
		"status":  "Error",
		"message": msg,
		"data":    nil,
		"error":   nil,
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
