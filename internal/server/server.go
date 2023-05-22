package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)

	StoreProcess(w http.ResponseWriter, r *http.Request)
	ReadProcess(w http.ResponseWriter, r *http.Request)
	ListProcesses(w http.ResponseWriter, r *http.Request)
	ListProcessesThreads(w http.ResponseWriter, r *http.Request)

	StoreThread(w http.ResponseWriter, r *http.Request)
	ReadThread(w http.ResponseWriter, r *http.Request)
	ListThreads(w http.ResponseWriter, r *http.Request)

	StoreQuant(w http.ResponseWriter, r *http.Request)
	ReadQuant(w http.ResponseWriter, r *http.Request)
	ListQuants(w http.ResponseWriter, r *http.Request)

	StoreSheet(w http.ResponseWriter, r *http.Request)
	ReadSheet(w http.ResponseWriter, r *http.Request)
	StoreRow(w http.ResponseWriter, r *http.Request)
}

func New(h Handler) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Logger)

	r.Post("/login", h.Login)
	r.Post("/register", h.Register)

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(jwt.TokenAuth))
		//r.Use(jwtauth.Authenticator)

		r.Post("/list/processes", h.ListProcesses)
		r.Get("/processes/{process}", h.ReadProcess)
		r.Get("/process/{process}/threads", h.ListProcessesThreads)
		r.Post("/processes", h.StoreProcess)

		r.Post("/list/threads", h.ListThreads)
		r.Get("/processes/{process}/{thread}", h.ReadThread)
		r.Post("/threads", h.StoreThread)

		r.Post("/quants", h.StoreQuant)
		r.Get("/processes/{process}/{thread}/{quant}", h.ReadQuant)
		r.Post("/list/quants", h.ListQuants)

		r.Post("/processes/{process}/sheet", h.StoreSheet)
		r.Get("/processes/{process}/sheet", h.ReadSheet)

		r.Post("/processes/{process}/sheet/row", h.StoreRow)

	})
	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
