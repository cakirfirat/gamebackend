package main

import (
	. "gamebackend/handlers"
	// . "gamebackend/helpers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()
	r.HandleFunc("/api/genesis", Genesis).Methods("POST")
	r.HandleFunc("/api/verify-code", VerifyCode).Methods("POST")
	server := &http.Server{
		Addr:    ":8096",
		Handler: r,
	}
	server.ListenAndServe()
}
