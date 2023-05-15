package handlers

import (
	"encoding/json"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
	"time"
)

func GenesisHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)
	locale := r.Header.Get("Accept-Language")

	user.CreatedAt = time.Now().String()
	user.UpdatedAt = time.Now().String()

	if phone, ok := jsonData["phoneNumber"].(string); ok {

		exists, _ := CheckPhoneNumber(phone)
		if !exists {
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

func VerifyCodeHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)

	if phone, ok := jsonData["phoneNumber"].(string); ok {

		user.PhoneNumber = phone

	} else {
		// PhoneNumber alanı yoksa veya türü string değilse hata döndür
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Telefon numarası eksik"))
		return
	}
	if code, ok := jsonData["code"].(string); ok {
		user.VerifyCode = code
	} else {
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

func UserHandler(w http.ResponseWriter, r *http.Request) {
	// var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)

	phoneNumber, err := GetJSONField(jsonData, "phoneNumber")
	identityNumber, err := GetJSONField(jsonData, "identityNumber")
	passportNumber, err := GetJSONField(jsonData, "passportNumber")
	email, err := GetJSONField(jsonData, "email")
	password, err := GetJSONField(jsonData, "password")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır. "))
		return
	}

	updateFields := map[string]interface{}{
		"PhoneNumber": phoneNumber,
		"NationalId":  identityNumber,
		"PassportId":  passportNumber,
		"Email":       email,
		"Password":    Md5Hash(password),
	}

	exists, id := CheckPhoneNumber(phoneNumber)
	if !exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Böyle bir kullanıcı mevcut değil."))
		return
	}

	if !UpdateUserFromPhone(phoneNumber, updateFields) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Güncellenirken bir hata oluştu. "))
		return
	}
	token, err := CreateJwt(id)

	response := map[string]interface{}{
		"accessToken":  token,
		"refreshToken": token,
	}
	responseJson, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}

func PatientHandler(w http.ResponseWriter, r *http.Request) {

	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)
	token := r.Header.Get("Authorization")

	firstName, err := GetJSONField(jsonData, "firstName")
	lastName, err := GetJSONField(jsonData, "lastName")
	dateOfBirth, err := GetJSONField(jsonData, "dateOfBirth")
	gender, err := GetJSONField(jsonData, "gender")
	identityNumber, err := GetJSONField(jsonData, "identityNumber")
	passportNumber, err := GetJSONField(jsonData, "passportNumber")
	nationalityCode, err := GetJSONField(jsonData, "nationalityCode")
	// hasMailActivation, err := GetJSONField(jsonData, "hasMailActivation")
	// careProviderId, err := GetJSONField(jsonData, "careProviderId")
	// patientType, err := GetJSONField(jsonData, "patientType")
	contactInfo, err := GetJSONFieldFromJson(jsonData, "contactInfo")
	// patientID, err := GetJSONField(jsonContactInfo,"patientID")
	phoneNumber, err := GetJSONField(contactInfo, "phoneNumber")
	emailAddress, err := GetJSONField(contactInfo, "emailAddress")
	// address, err := GetJSONField(jsonContactInfo,"address")
	// addressLongitude, err := GetJSONField(jsonContactInfo,"addressLongitude")
	// addressLatitude, err := GetJSONField(jsonContactInfo,"addressLatitude")
	// relativeFullName, err := GetJSONField(jsonContactInfo,"relativeFullName")
	// relativePhoneNumber, err := GetJSONField(jsonContactInfo,"relativePhoneNumber")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır. "))
		CheckError(err)
		return
	}

	updateFields := map[string]interface{}{
		"PhoneNumber": phoneNumber,
		"NationalId":  identityNumber,
		"PassportId":  passportNumber,
		"Email":       emailAddress,
		"Name":        firstName,
		"LastName":    lastName,
		"BirthDate":   dateOfBirth,
		"Gender":      gender,
		"Nationality": nationalityCode,
		"IsPatient":   1,
	}

	id, err := ExtractUserId(token)

	user, err := GetUserById(id)

	if user == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Böyle bir kullanıcı bulunmamaktadır."))
		CheckError(err)
		return
	}

	if !UpdateUserFromId(id, updateFields) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Güncelleme hatası."))
		CheckError(err)
		return
	}

	response := map[string]interface{}{
		"UserId": id,
	}

	responseJson, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// var user User
	var jsonData map[string]interface{}
	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)
	CheckError(errorDecoder)

	phoneNumber, err := GetJSONField(jsonData, "phoneNumber")
	password, err := GetJSONField(jsonData, "password")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır. "))
		return
	}

	validateUser, err := CheckPhoneNumberAndPassword(phoneNumber, Md5Hash(password))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Kullanıcı bilgileri yanlış. "))
		return
	}

	if validateUser == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Kullanıcı adı veya parola hatalı."))
		return
	}

	token, err := CreateJwt(validateUser.Id)

	name := validateUser.Name
	lastName := validateUser.LastName
	email := validateUser.Email
	birthDate := validateUser.BirthDate
	gender := validateUser.Gender

	user := map[string]interface{}{
		"name":      name,
		"lastName":  lastName,
		"email":     email,
		"birthDate": birthDate,
		"gender":    gender,
	}

	response := map[string]interface{}{
		"accessToken":  token,
		"refreshToken": token,
		"user":         user,
	}
	responseJson, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}
