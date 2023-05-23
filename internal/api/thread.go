package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/thread"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// Store process route that accepts json representation of Process and stores it to Storage
func (h *Handler) StoreThread(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t thread.Thread
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.StoreThread(usr, t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ListThreads(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var t thread.Thread
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	threads, err := h.S.ListThreads(usr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err = json.Marshal(threads)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) ReadThread(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	threadTitle := chi.URLParam(r, "thread")
	thread, err := thread.New(procTitle, threadTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.S.ReadThread(usr, thread)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)

}
