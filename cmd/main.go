package main

import (
	"github.com/gorilla/mux"
	"github.com/tinarao/url-shortener-go/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", handlers.Shortener).Methods(http.MethodPost)
	http.Handle("/", router)

	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal(err)
	}
}
