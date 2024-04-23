package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tinarao/url-shortener-go/db"
	"github.com/tinarao/url-shortener-go/handlers"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", handlers.Shortener).Methods("POST")
	router.HandleFunc("/all", handlers.GetAllLinks).Methods("GET")
	router.HandleFunc("/m/{linkID}", handlers.RedirectToShortened).Methods("GET")

	db.Connect()

	fmt.Println("Running and serving at 3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
