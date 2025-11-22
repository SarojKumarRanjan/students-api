package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/SarojKumarRanjan/students-api/internal/types"
	"github.com/SarojKumarRanjan/students-api/internal/utils/response"
	"github.com/go-playground/validator"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// now validate the request

		if err := validator.New().Struct(student); err != nil {
			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateError))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"message": "ok"})
	}
}
