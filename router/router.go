package router

import (
	"api-boilerplate/handler"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// NewRouter func
func NewRouter() *chi.Mux {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/booking", h.StoreBooking)

	return r
}
