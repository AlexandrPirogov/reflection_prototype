package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflection_prototype/internal/core/quant"
)

// Store process route that accepts json representation of Process and stores it to Storage
func (h *Handler) StoreQuant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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

	err = h.S.StoreQuant(q)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) ReadQuants(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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

	quants, err := h.S.ReadQuants(q)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err = json.Marshal(quants)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
