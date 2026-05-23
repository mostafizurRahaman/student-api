package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) {

	// ! Header setting:
	w.Header().Set("Content-Type", "application/json")
	// ! Set status
	w.WriteHeader(status)

	// ! Send the encoded response data:
	json.NewEncoder(w).Encode(data)

}
