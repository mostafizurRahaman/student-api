package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

//  ?? Validation error :

func ValidationError(errs validator.ValidationErrors) ErrorResponse {

	var errMsgs []string

	for _, err := range errs {
		fmt.Println(err.ActualTag())
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s field is required", err.Field()))
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("%s should be min %s character.", err.Field(), err.Param()))
		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("%s should be max %s character.", err.Field(), err.Param()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("Invalid  field %s: Error: %s", err.Field(), err.Error()))
		}
	}

	return ErrorResponse{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}

}
