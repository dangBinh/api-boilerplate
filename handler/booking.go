package handler

import (
	"api-boilerplate/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// BookingHandler ...
type BookingHandler interface {
	CreateBooking(w http.ResponseWriter, r *http.Request)
	UpdateBooking(w http.ResponseWriter, r *http.Request)
	DetailBooking(w http.ResponseWriter, r *http.Request)
}

type bookingHandlerImpl struct {
	bookingDAO storage.BookingStorage
}

// CreateBooking ...
func (h *bookingHandlerImpl) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking storage.Booking
	getBody(&booking, r.Body)

	h.bookingDAO.Create(&booking)
	w.WriteHeader(200)
	w.Write([]byte("Successfully create booking"))
}

// UpdateBooking ...
func (h *bookingHandlerImpl) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	var booking storage.Booking
	getBody(&booking, r.Body)

	h.bookingDAO.Update(&booking)
	w.WriteHeader(200)
	w.Write([]byte("Successfully update booking"))
}

func (h *bookingHandlerImpl) DetailBooking(w http.ResponseWriter, r *http.Request) {
	strID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	booking, _ := h.bookingDAO.ByID(id)
	fmt.Printf("%v\n", booking)
	res, err := json.Marshal(booking)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Write([]byte(res))
}
