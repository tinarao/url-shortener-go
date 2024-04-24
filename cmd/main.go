package main

import (
	"github.com/gorilla/mux"
	"github.com/tinarao/url-shortener-go/db"
	"github.com/tinarao/url-shortener-go/handlers"
	"github.com/tinarao/url-shortener-go/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/l/{alias}", handlers.RedirectToShortened).Methods("GET")
	r.HandleFunc("/shorten", handlers.Shortener).Methods("POST")
	r.HandleFunc("/get-all", handlers.GetAllLinks).Methods("GET")
	r.HandleFunc("/get-one/{alias}", handlers.GetByAlias).Methods("GET")
	r.HandleFunc("/delete/all", handlers.DeleteAllLinks).Methods("DELETE")

	db.Connect()

	slog.Info("Running and serving at 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		slog.Error("Error while starting a server", err)
		os.Exit(1)
	}
}
