package models

// import (
// 	. "gamebackend/helpers"
// )

type Game struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Detail    string `json:"Detail"`
	PhotoUrl  string `json:"PhotoUrl"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func InsertGame(game Game) (string, error) {
	sqlQuery := `
        INSERT INTO game (
            Name, Detail, PhotoUrl, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?)
        RETURNING id
    `

	var id string
	err := db.QueryRow(sqlQuery, game.Name, game.Detail, game.PhotoUrl, game.CreatedAt, game.UpdatedAt).Scan(&id)
	if err != nil {
		return "0", err
	}

	return id, nil
}
