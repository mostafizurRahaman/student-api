package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mostafizurRahaman/student-api/internal/types"
	response "github.com/mostafizurRahaman/student-api/internal/utils"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Decode the request:

		// ? declare a struct for response decode

		var student types.Student

		// ? Decode here and store into student variable
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadGateway, response.ErrorResponse(fmt.Errorf("empty error")))
			return
		}

		// ? If error is EOF and not nil :
		if err != nil {
			response.WriteJson(w, http.StatusBadGateway, response.ErrorResponse(err))
			return
		}

		response.WriteJson(w, http.StatusBadGateway, map[string]string{"status": "Ok"})

	}

}
