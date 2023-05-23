package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/process"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// Store process route that accepts json representation of Process and stores it to Storage
func (h *Handler) StoreProcess(w http.ResponseWriter, r *http.Request) {
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

	var proc process.Process
	err = json.Unmarshal(body, &proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.StoreProcess(usr, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ListProcesses(w http.ResponseWriter, r *http.Request) {
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

	var proc process.Process
	err = json.Unmarshal(body, &proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	processes, err := h.S.ListProcesses(usr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err = json.Marshal(processes)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) ReadProcess(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	process, err := process.New(procTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.S.ReadProcess(usr, process)
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

func (h *Handler) ListProcessesThreads(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	process, err := process.New(procTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	threads, err := h.S.ListProcessesThreads(usr, process)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(threads)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
