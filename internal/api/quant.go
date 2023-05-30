package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/quant"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5"
)

// Store process route that accepts json representation of Process and stores it to Storage
func (h *Handler) StoreQuant(w http.ResponseWriter, r *http.Request) {
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

	var q quant.Quant
	err = json.Unmarshal(body, &q)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.StoreQuant(usr, q)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ListQuants(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	quants, err := h.S.ListQuants(usr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(quants)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) ReadQuant(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	procTitle := chi.URLParam(r, "process")
	threadTitle := chi.URLParam(r, "thread")
	quantTitle := chi.URLParam(r, "quant")

	quant, err := quant.New(procTitle, threadTitle, quantTitle, "")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.S.ReadQuant(usr, quant)
	if errors.Is(err, pgx.ErrNoRows) {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
