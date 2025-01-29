package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string
	Error  string
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: "400",
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("filed %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("filed %s is invalid", err.Field()))
		}
	}
	return Response{
		Status: "500",
		Error:strings.Join(errMsgs,", "),
	}
}
