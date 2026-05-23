package customers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mostafizurRahaman/customer-api/internal/types"
	"github.com/mostafizurRahaman/customer-api/utils/response"
)

// ! customer create handler
func New() http.HandlerFunc {

	// ! return an handler function
	return func(w http.ResponseWriter, r *http.Request) {

		// ? prepare an struct:
		var customer types.Customer

		err := json.NewDecoder(r.Body).Decode(&customer)

		// ?. Parse the json:
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Println(customer)

		response.WriteJson(w, http.StatusOK, customer)

	}

}
