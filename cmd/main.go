package main

import (
	"maratproject/pkg/api"
	"maratproject/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.New("postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe("localhost:8090"))
}
