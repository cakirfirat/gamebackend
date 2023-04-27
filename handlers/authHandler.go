package handlers

import (
	"encoding/json"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
	"time"
)

func Genesis(w http.ResponseWriter, r *http.Request) {
	var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)
	locale := r.Header.Get("Accept-Language")
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if phone, ok := jsonData["PhoneNumber"].(string); ok {

		if CheckPhoneNumber(phone) {
			user.PhoneNumber = phone
			otp := CreateOtp()
			user.VerifyCode = otp
			SendSms(phone, Localizate(locale, "OTP Message")+otp)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Bu telefon numarası veritabanında kayıtlıdır."))
			return
		}

	} else {
		// PhoneNumber alanı yoksa veya türü string değilse hata döndür
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Yanlış veya eksik bir bilgi girdiniz."))
		return
	}

	userUUID, err := GenerateUUID()
	if err != nil {
		// UUID oluşturma hatasıyla ilgili uygun hata mesajını döndürün
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("UUID oluşturulamadı."))
		return
	}
	user.Id = userUUID
	if InsertUser(user) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Kayıt ekleme başarılı."))
		return
	}

}

func VerifyCode(w http.ResponseWriter, r *http.Request) {
	var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)

	if phone, ok := jsonData["PhoneNumber"].(string); ok {

		user.PhoneNumber = phone

	} else {
		// PhoneNumber alanı yoksa veya türü string değilse hata döndür
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Telefon numarası eksik"))
		return
	}
	if code, ok := jsonData["Code"].(string); ok {
		user.VerifyCode = code
	} else {
		// PhoneNumber alanı yoksa veya türü string değilse hata döndür
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doğrulama kodu eksik"))
		return
	}

	if CheckVerifyCode(user.PhoneNumber, user.VerifyCode) {

		updateFields := map[string]interface{}{
			"IsVerify": 1,
		}
		if UpdateUserFromPhone(user.PhoneNumber, updateFields) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("İşlem başarılı"))
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bir hata oluştu"))
			return
		}

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Doğrulama kodunuz yanlış"))
		return
	}

}
