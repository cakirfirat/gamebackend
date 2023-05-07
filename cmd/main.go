package main

import (
	. "gamebackend/handlers"
	. "gamebackend/helpers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()
	r.HandleFunc("/api/genesis", GenesisHandler).Methods("POST")
	r.HandleFunc("/api/verify-code", VerifyCodeHandler).Methods("POST")
	r.HandleFunc("/api/user", UserHandler).Methods("POST")
	r.HandleFunc("/api/login", LoginHandler).Methods("POST")
	r.Handle("/api/patient-user", ValidateJwt(PatientHandler)).Methods("POST")

	r.Handle("/api/setgame", ValidateJwt(SetGameHandler)).Methods("POST")
	server := &http.Server{
		Addr:    ":8096",
		Handler: r,
	}
	server.ListenAndServe()
}
