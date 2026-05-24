package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	// ! Header setting:
	w.Header().Set("Content-Type", "application/json")
	// ! Set status
	w.WriteHeader(status)

	// ! Send the encoded response data:
	return json.NewEncoder(w).Encode(data)

}

type ErrorStatus string

const (
	StatusOk    ErrorStatus = "OK"
	StatusError             = "Error"
)

// General error type :

type Response struct {
	Status ErrorStatus `json:"status"`
	Error  string      `json:"message"`
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}
