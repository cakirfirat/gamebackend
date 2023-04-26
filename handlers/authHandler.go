package handlers

import (
	"encoding/json"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
)

func Genesis(w http.ResponseWriter, r *http.Request) {
	var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)
	locale := r.Header.Get("Accept-Language")

	if phone, ok := jsonData["PhoneNumber"].(string); ok {
		user.PhoneNumber = phone
		SendSms(phone, Localizate(locale, "OTP Message")+CreateOtp())
	} else {
		// PhoneNumber alanı yoksa veya türü string değilse hata döndür
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Yanlış veya eksik bir bilgi girdiniz"))
		return
	}
	if InsertUser(user) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Kayıt ekleme başarılı."))
		return
	}

}
