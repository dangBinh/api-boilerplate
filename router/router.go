package router

import (
	"api-boilerplate/handler"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// NewRouter func
func NewRouter() *chi.Mux {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/booking", h.CreateBooking)
	r.Put("/booking/{id}", h.UpdateBooking)

	r.Post("/customer", h.CreateCustomer)
	return r
}
