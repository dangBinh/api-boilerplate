package handler

import (
	"api-boilerplate/storage"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CustomerHandler ...
type CustomerHandler interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
}

type customerHandlerImpl struct {
	customerDAO storage.CustomerStorage
}

func (h *customerHandlerImpl) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var customer storage.Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		panic(err)
	}

	h.customerDAO.Create(&customer)
	w.WriteHeader(200)
	w.Write([]byte("Successfully create customer"))
}
