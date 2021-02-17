package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mutant-app/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/mutant", handlers.Mutant).Methods("POST")
	r.HandleFunc("/stats", handlers.Stats).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening...")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
