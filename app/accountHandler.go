package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-hexagonal-arch/dto"
	"github.com/go-hexagonal-arch/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	var request dto.NewAccountRequest
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customer_id
		acc, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, acc)
		}
	}
}
