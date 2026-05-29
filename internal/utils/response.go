package response

import (
	"encoding/json"
	"net/http"
)

// ?? Create a struct for response
type ErrorResponse struct {
	Status StatusType `json:"status"`
	Error  string     `json:"error"`
}

type StatusType string

const (
	StatusSuccess StatusType = "OK"
	StatusError   StatusType = "Error"
)

// ?? Write json Utils function:
func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	// Encode the code return response:
	// ? Write the header type :
	w.Header().Set("Content-Type", "application/json")

	// ? Update status code :
	w.WriteHeader(status)

	// ? Now Encode the json response:
	return json.NewEncoder(w).Encode(data)

}

// ?? Create a generic for response type:

func GeneralError(err error) ErrorResponse {

	return ErrorResponse{
		Status: StatusError,
		Error:  err.Error(),
	}
}
