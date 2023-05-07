package handlers

import (
	"encoding/json"
	. "gamebackend/helpers"
	. "gamebackend/models"
	"net/http"
	"time"
)

func SetGameHandler(w http.ResponseWriter, r *http.Request) {
	var game Game
	var jsonData map[string]interface{}

	userUUID, err := GenerateUUID()

	errorDecoder := json.NewDecoder(r.Body).Decode(&jsonData)

	CheckError(errorDecoder)

	name, err := GetJSONField(jsonData, "Name")
	detail, err := GetJSONField(jsonData, "Detail")

	game.Id = userUUID
	game.Name = name
	game.Detail = detail
	game.CreatedAt = time.Now().String()
	game.UpdatedAt = time.Now().String()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bilgiler dolu olmalıdır."))
		return
	}

	id, err := InsertGame(game)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Oyun eklenirken bir hata oluştu."))
		return
	}

	updateFields := map[string]interface{}{
		"gameId": id,
		"name":   game.Name,
		"detail": game.Detail,
	}

	responseJson, err := json.Marshal(updateFields)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
	CheckError(err)
	return

}
