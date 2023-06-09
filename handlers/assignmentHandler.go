package handlers

import (
	"encoding/json"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
	"time"
)

func AssignPatient(w http.ResponseWriter, r *http.Request) {

	var assignment Assignment
	var jsonData map[string]interface{}

	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)

	CheckError(errorDecoder)

	doctorId, err := GetJSONField(jsonData, "doctorId")
	userId, err := GetJSONField(jsonData, "userId")

	assignment.DoctorId = doctorId
	assignment.UserId = userId
	assignment.IsDeleted = "0"
	assignment.CreatedAt = time.Now().String()
	assignment.UpdatedAt = time.Now().String()
	CheckError(err)

	if !SetPatientAssignment(assignment) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Atama esnasında bir hata oluştu."))
		return
	}

	assignmentData := map[string]interface{}{
		"doctorId":  doctorId,
		"userId":    userId,
		"createdAt": assignment.CreatedAt,
		"updatedAt": assignment.UpdatedAt,
	}
	jsonAssignmentData, err := json.Marshal(assignmentData)

	CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonAssignmentData)
	return

}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	id, err := ExtractUserId(token)
	if err != nil {
		// Hata durumunu kontrol etme
		// ...
	}

	assignments, err := GetPatientAssignments(id)
	if err != nil {
		// Hata durumunu kontrol etme
		// ...
	}

	var assignedUsers []User
	for _, assignment := range assignments {
		users, err := GetUserByAssignment(assignment.UserId)
		if err != nil {
			// Hata durumunu kontrol etme
			// ...
		}
		assignedUsers = append(assignedUsers, users...)
	}

	responseJson, err := json.Marshal(assignedUsers)
	if err != nil {
		// Hata durumunu kontrol etme
		// ...
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
