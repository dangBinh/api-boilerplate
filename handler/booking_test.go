package handler

import (
	"booking/storage"
	"testing"
)

func TestCreateBooking(t *testing.T) {
	h := &handlerImpl{
		bookingDAO: &storage.MockBookingStorage{},
	}
	h.StoreBooking()
}
