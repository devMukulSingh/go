package hello

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"tutorial/internal/storage"
	"tutorial/internal/types"
	"tutorial/internal/utils/response"

	"github.com/go-playground/validator/v10"
)

func GetHello() http.HandlerFunc {
	return func( w http.ResponseWriter, r * http.Request){
		w.Write([]byte("Hello world"))
	}
}

func PostHello(storage storage.Storage) http.HandlerFunc{
	return func( w http.ResponseWriter, r * http.Request){

		var postData types.PostData

		err := json.NewDecoder(r.Body).Decode(&postData); 

		if errors.Is(err,io.EOF){
			response.WriteJson(w,http.StatusBadRequest,err.Error())
			return;
		}

		if err!= nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError((err)))
			return
		}

		//request data validation
		if err := validator.New().Struct(postData); err!=nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return;
		}

		lastId,err := storage.CreateStudent(
			postData.Email,
			postData.Name,
			postData.Id,
		)

		if err!= nil{
			response.WriteJson(w,http.StatusInternalServerError,err)
			return
		}
		
		response.WriteJson(w, http.StatusCreated,lastId )
	}
}