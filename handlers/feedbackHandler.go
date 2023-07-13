package handlers

import (
	"encoding/json"
	"fmt"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
)

func SetFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var feedback Feedback
	var jsonData map[string]interface{}

	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)

	CheckError(errorDecoder)

	comment, err := GetJSONField(jsonData, "comment")

	feedback.Comment = comment
	feedback.Ip = "ip"
	feedback.Mac = "mac"

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır."))
		return
	}

	id, err := InsertFeedback(feedback)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Oyun eklenirken bir hata oluştu."))
		return
	}

	updateFields := map[string]interface{}{
		"scoreId": id,
	}

	responseJson, err := json.Marshal(updateFields)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
