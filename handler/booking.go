package handler

import (
	"api-boilerplate/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BookingHandler ...
type BookingHandler interface {
	CreateBooking(w http.ResponseWriter, r *http.Request)
	UpdateBooking(w http.ResponseWriter, r *http.Request)
}

type bookingHandlerImpl struct {
	bookingDAO storage.BookingStorage
}

// CreateBooking ...
func (h *bookingHandlerImpl) CreateBooking(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var booking storage.Booking
	err = json.Unmarshal(body, &booking)
	if err != nil {
		panic(err)
	}

	h.bookingDAO.Create(&booking)
	w.WriteHeader(200)
	w.Write([]byte("Successfully create booking"))
}

// UpdateBooking ...
func (h *bookingHandlerImpl) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var booking storage.Booking
	err = json.Unmarshal(body, &booking)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", booking)
	h.bookingDAO.Update(&booking)
	w.WriteHeader(200)
	w.Write([]byte("Successfully update booking"))
}
