package handlers

import "net/http"

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello motherfucker deniz cem yıldız"))
}
