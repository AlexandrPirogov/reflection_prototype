package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (h *Handler) StoreSheet(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

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

	err = h.S.StoreSheet(usr, shet, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ReadSheet(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	proc, err := process.New(procTitle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shet, _ := h.S.ReadSheet(usr, proc)

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
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	proc, _ := process.New(procTitle)

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

	err = h.S.AddRow(usr, row, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) MarkRow(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	proc, _ := process.New(procTitle)

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

	err = h.S.MarkRow(usr, row, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) StartWork(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	proc, _ := process.New(procTitle)

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

	err = h.S.StartWork(usr, row, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handler) StopWork(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	proc, _ := process.New(procTitle)

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

	err = h.S.StopWork(usr, row, proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
