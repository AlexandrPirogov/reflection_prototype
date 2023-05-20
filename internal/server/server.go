package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler interface {
	StoreProcess(w http.ResponseWriter, r *http.Request)
	ReadProcesses(w http.ResponseWriter, r *http.Request)

	StoreThread(w http.ResponseWriter, r *http.Request)
	ReadThreads(w http.ResponseWriter, r *http.Request)

	StoreQuant(w http.ResponseWriter, r *http.Request)
	ReadQuants(w http.ResponseWriter, r *http.Request)
}

func New(h Handler) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)
	r.Post("/list/processes", h.ReadProcesses)
	r.Post("/processes", h.StoreProcess)

	r.Post("/list/threads", h.ReadThreads)
	r.Post("/threads", h.StoreThread)

	r.Post("/quants", h.StoreQuant)
	r.Post("/list/quants", h.ReadQuants)
	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
