package handler

import (
	"api-boilerplate/storage"
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
