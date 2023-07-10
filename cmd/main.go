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

	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/api/genesis", GenesisHandler).Methods("POST")
	r.HandleFunc("/api/verify-code", VerifyCodeHandler).Methods("POST")
	r.HandleFunc("/api/user", UserHandler).Methods("POST")
	r.HandleFunc("/api/login", LoginHandler).Methods("POST")
	r.Handle("/api/patient-user", ValidateJwt(PatientHandler)).Methods("POST")

	r.Handle("/api/add-game", ValidateJwt(SetGameHandler)).Methods("POST")
	r.Handle("/api/games", ValidateJwt(GetGamesHandler)).Methods("GET")
	r.Handle("/api/get-user-games", ValidateJwt(GetGamesByUser)).Methods("POST")

	r.Handle("/api/add-score", ValidateJwt(SetScoreHandler)).Methods("POST")
	r.Handle("/api/get-score", ValidateJwt(GetScoreHandler)).Methods("POST")

	r.Handle("/api/assign-patient", ValidateJwt(AssignPatient)).Methods("POST")
	r.Handle("/api/assignments", ValidateJwt(GetAssignment)).Methods("GET")

	/* Alakasız endpoint */
	r.HandleFunc("/api/file-upload", FileUploadHandler).Methods("POST")
	r.HandleFunc("/api/feedback", SetFeedbackHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8096",
		Handler: handlers.CORS()(r),
	}
	server.ListenAndServe()
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
