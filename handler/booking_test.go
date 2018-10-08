package handler

import (
	"api-boilerplate/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBooking(t *testing.T) {
	r, err := http.NewRequest("POST", "/booking", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := handlerImpl{
		&bookingHandlerImpl{bookingDAO: storage.NewBookingStorage()},
		nil,
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateBooking)

	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Successfully create booking`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
