package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflection_prototype/internal/core/auth/jwt"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/report"
	"reflection_prototype/internal/core/sheet"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
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

	var report report.Report
	err = json.Unmarshal(body, &report)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.CreateReport(usr, report)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *Handler) FillReport(w http.ResponseWriter, r *http.Request) {
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

	reportTitle := chi.URLParam(r, "report")
	rep := report.New(reportTitle)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var sheetRow sheet.SheetRow
	err = json.Unmarshal(body, &sheetRow)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.S.FillReport(usr, sheetRow, process, rep)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ViewReport(w http.ResponseWriter, r *http.Request) {
	usr, err := jwt.UserFromToken(jwtauth.TokenFromHeader(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(401)
		return
	}

	reportTitle := chi.URLParam(r, "report")
	rep := report.New(reportTitle)

	res, err := h.S.ReadReport(usr, rep)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
