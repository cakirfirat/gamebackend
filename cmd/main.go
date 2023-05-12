package main

import (
	. "gamebackend/handlers"
	. "gamebackend/helpers"

	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting...")

	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/api/genesis", GenesisHandler).Methods("POST")
	r.HandleFunc("/api/verify-code", VerifyCodeHandler).Methods("POST")
	r.HandleFunc("/api/user", UserHandler).Methods("POST")
	r.HandleFunc("/api/login", LoginHandler).Methods("POST")
	r.Handle("/api/patient-user", ValidateJwt(PatientHandler)).Methods("POST")

	r.Handle("/api/add-game", ValidateJwt(SetGameHandler)).Methods("POST")
	r.Handle("/api/games", ValidateJwt(GetGameHandler)).Methods("GET")

	r.Handle("/api/add-score", ValidateJwt(SetScoreHandler)).Methods("POST")
	server := &http.Server{
		Addr:    ":8096",
		Handler: handlers.CORS(headers, methods, origins)(r),
	}
	server.ListenAndServe()
}
