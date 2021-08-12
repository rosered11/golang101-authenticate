package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rosered11/golang101-authenticate/dto"
	"github.com/rosered11/golang101-authenticate/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ch *AccountHandlers) saveNewAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customer_id
		response, appErr := ch.service.NewAccount(request)
		if appErr != nil {
			writeResponse(rw, appErr.Code, appErr.AsMessage())
		} else {
			writeResponse(rw, http.StatusCreated, response)
		}
	}
}
