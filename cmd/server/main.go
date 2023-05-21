package main

import (
	"log"
	"reflection_prototype/internal/api"
	"reflection_prototype/internal/server"
	storage "reflection_prototype/internal/storage/postgres"
)

func main() {
	h := api.Handler{
		S: storage.New(),
	}
	s := server.New(&h)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
