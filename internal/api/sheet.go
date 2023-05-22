package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) StoreSheet(w http.ResponseWriter, r *http.Request) {
	procTitle := chi.URLParam(r, "process")
	proc, err := process.New(procTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var shet sheet.Sheet
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &shet)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.StoreSheet(shet, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) ReadSheet(w http.ResponseWriter, r *http.Request) {
	procTitle := chi.URLParam(r, "process")
	proc, err := process.New(procTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shet, err := h.S.ReadSheet(proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(shet)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) StoreRow(w http.ResponseWriter, r *http.Request) {
	procTitle := chi.URLParam(r, "process")
	shet := sheet.New(procTitle, "")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var row sheet.SheetRow
	err = json.Unmarshal(body, &row)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.AddRow(row, shet)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
