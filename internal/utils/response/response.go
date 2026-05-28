package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

// Generic response type:

const (
	StatusOk    = "Success"
	StatusError = "Error"
)

type ErrorResponse struct {
	Status  string
	Success bool
	Error   string
}

func Response(err error) ErrorResponse {

	return ErrorResponse{
		Status:  StatusOk,
		Success: false,
		Error:   err.Error(),
	}
}
