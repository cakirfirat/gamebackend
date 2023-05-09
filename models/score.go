package models

// . "gamebackend/helpers"

type Score struct {
	Id        string `json:"Id"`
	GameId    string `json:"GameId"`
	UserId    string `json:"UserId"`
	Correct   string `json:"Correct"`
	Wrong     string `json:"Wrong"`
	Total     string `json:"Total"`
	Speed     string `json:"Speed"`
	PlayTime  string `json:"PlayTime"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func InsertScore(score Score) (string, error) {
	sqlQuery := `
        INSERT INTO score (
            Id, GameId, UserId, Correct, Wrong, Total, Speed,
			PlayTime, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?,?,?,?,?,?)
    `

	_, err := db.Exec(sqlQuery, score.Id, score.GameId, score.UserId, score.Correct, score.Wrong, score.Total, score.Speed,
		score.PlayTime, score.CreatedAt, score.UpdatedAt)
	if err != nil {
		return "0", err
	}

	return score.Id, nil
}

func GetScore(gameId, userId string) ([]Score, error) {
	var scores []Score

	sqlQuery := `
        SELECT * FROM score
        WHERE GameId = ? AND UserId = ?
    `
	rows, err := db.Query(sqlQuery, gameId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var score Score
		err := rows.Scan(&score.Id, &score.GameId, &score.UserId, &score.Correct, &score.Wrong, &score.Total, &score.Speed, &score.PlayTime, &score.CreatedAt, &score.UpdatedAt)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scores, nil
}
