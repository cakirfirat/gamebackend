package models

// import (
// 	. "gamebackend/helpers"
// )

type Game struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Detail    string `json:"Detail"`
	Tag       string `json:"Tag"`
	PhotoUrl  string `json:"PhotoUrl"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func InsertGame(game Game) (string, error) {
	sqlQuery := `
        INSERT INTO game (
            Id, Name, Detail, Tag, PhotoUrl, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?,?,?)
    `

	_, err := db.Exec(sqlQuery, game.Id, game.Name, game.Detail, game.Tag, game.PhotoUrl, game.CreatedAt, game.UpdatedAt)
	if err != nil {
		return "0", err
	}

	return game.Id, nil
}

func GetGames() ([]Game, error) {
	var games []Game

	rows, err := db.Query("SELECT * FROM game")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var game Game
		err := rows.Scan(&game.Id, &game.Name, &game.Detail, &game.Tag, &game.PhotoUrl, &game.CreatedAt, &game.UpdatedAt)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return games, nil
}
