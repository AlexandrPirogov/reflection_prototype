package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflection_prototype/internal/core/thread"

	"github.com/go-chi/chi/v5"
)

// Store process route that accepts json representation of Process and stores it to Storage
func (h *Handler) StoreThread(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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

	err = h.S.StoreThread(t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ListThreads(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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

	threads, err := h.S.ListThreads()
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
	procTitle := chi.URLParam(r, "process")
	threadTitle := chi.URLParam(r, "thread")
	thread, err := thread.New(procTitle, threadTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.S.ReadThread(thread)
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
