package handler

import (
	"api-boilerplate/storage"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	StoreBooking(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	bookingDAO storage.BookingStorage
}

func NewHandler() Handler {
	return &handlerImpl{bookingDAO: storage.NewBookingStorage()}
}

// StoreBooking ...
func (h *handlerImpl) StoreBooking(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var booking storage.Booking
	err = json.Unmarshal(body, &booking)
	if err != nil {
		panic(err)
	}

	// err = h.bookingDAO.CreateBooking(&booking)
	h.bookingDAO.Create(&booking)
	fmt.Println(err)
}
