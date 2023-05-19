package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflection_prototype/internal/core/process"
)

func (h *Handler) StoreProcess(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
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

	err = h.S.StoreProcess(proc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
