package handler

import (
	"api-boilerplate/storage"
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
	var customer storage.Customer
	getBody(customer, r.Body)

	h.customerDAO.Create(&customer)
	w.WriteHeader(200)
	w.Write([]byte("Successfully create customer"))
}

func (h *customerHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {

}
