package handlers

import (
	"encoding/json"
	"fmt"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
	"time"
)

func SetScoreHandler(w http.ResponseWriter, r *http.Request) {
	var score Score
	var jsonData map[string]interface{}
	token := r.Header.Get("Authorization")
	userId, err := ExtractUserId(token)
	scoreUUID, err := GenerateUUID()

	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)

	CheckError(errorDecoder)

	gameId, err := GetJSONField(jsonData, "gameId")
	correct, err := GetJSONField(jsonData, "correct")
	wrong, err := GetJSONField(jsonData, "wrong")
	total, err := GetJSONField(jsonData, "total")
	speed, err := GetJSONField(jsonData, "speed")
	playTime, err := GetJSONField(jsonData, "playTime")

	score.Id = scoreUUID
	score.GameId = gameId
	score.UserId = userId
	score.Correct = correct
	score.Wrong = wrong
	score.Total = total
	score.Speed = speed
	score.PlayTime = playTime
	score.CreatedAt = time.Now().String()
	score.UpdatedAt = time.Now().String()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır."))
		return
	}

	id, err := InsertScore(score)
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

func GetScoreHandler(w http.ResponseWriter, r *http.Request) {
	var jsonData map[string]interface{}

	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)

	CheckError(errorDecoder)

	gameId, err := GetJSONField(jsonData, "gameId")
	id, err := GetJSONField(jsonData, "userId")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır."))
		return
	}

	score, err := GetScore(gameId, id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Skor bulunmamaktadır."))
		return
	}

	responseJson, err := json.Marshal(score)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}
