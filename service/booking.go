package service

import (
	"booking/storage"
)

// CreateBooking ...
func CreateBooking(booking *storage.Booking) (int, error) {
	var bookingStorage = storage.NewBookingStorage()
	err := bookingStorage.CreateBooking(booking)
	if err != nil {
		return 0, err
	}
	return booking.ID, nil
}
