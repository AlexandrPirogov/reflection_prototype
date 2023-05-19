package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler interface {
	StoreProcess(w http.ResponseWriter, r *http.Request)
}

func New(h Handler) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	r.Post("/processes/", h.StoreProcess)
	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
