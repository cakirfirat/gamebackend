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

type GameWithScores struct {
	Game   Game    `json:"Game"`
	Scores []Score `json:"Scores"`
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

func GetGamesWithScoresByUserId(userId string) ([]GameWithScores, error) {
	var gamesWithScores []GameWithScores

	query := `
        SELECT g.Id, g.Name, g.Detail, g.Tag, g.PhotoUrl, s.Correct, s.Wrong, s.Total, s.Speed, s.PlayTime
        FROM game g
        INNER JOIN score s ON g.Id = s.GameId
        WHERE s.UserId = ?
        ORDER BY g.Id
    `

	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var currentGameWithScores *GameWithScores

	for rows.Next() {
		var gameWithScores GameWithScores
		var game Game
		var score Score

		err := rows.Scan(
			&game.Id, &game.Name, &game.Detail, &game.Tag, &game.PhotoUrl,
			&score.Correct, &score.Wrong, &score.Total, &score.Speed, &score.PlayTime,
		)
		if err != nil {
			return nil, err
		}

		if currentGameWithScores == nil || currentGameWithScores.Game.Id != game.Id {
			// Yeni bir oyun başladı, GameWithScores yapısını oluştur
			gameWithScores.Game = game
			gameWithScores.Scores = []Score{score}
			gamesWithScores = append(gamesWithScores, gameWithScores)
			currentGameWithScores = &gamesWithScores[len(gamesWithScores)-1]
		} else {
			// Aynı oyunun skorlarını eklemeye devam et
			currentGameWithScores.Scores = append(currentGameWithScores.Scores, score)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gamesWithScores, nil
}
