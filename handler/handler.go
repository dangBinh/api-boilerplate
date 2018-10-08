package handler

import (
	"api-boilerplate/storage"
	"encoding/json"
	"io"
	"io/ioutil"
)

// Handler ...
type Handler interface {
	BookingHandler
	CustomerHandler
}

type handlerImpl struct {
	*bookingHandlerImpl
	*customerHandlerImpl
}

// NewHandler ...
func NewHandler() Handler {
	return handlerImpl{
		&bookingHandlerImpl{bookingDAO: storage.NewBookingStorage()},
		&customerHandlerImpl{customerDAO: storage.NewCustomerStorage()},
	}
}

func getBody(o interface{}, readClose io.ReadCloser) {
	body, err := ioutil.ReadAll(readClose)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		panic(err)
	}
}
