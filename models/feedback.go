package models

// . "gamebackend/helpers"

type Feedback struct {
	Comment string `json:"comment"`
	Ip      string `json:"ip"`
	Mac     string `json:"mac"`
}

func InsertFeedback(feedback Feedback) (string, error) {
	sqlQuery := `
        INSERT INTO feedback (
            Comment, Ip, Mac
        ) VALUES (?,?,?)
    `

	_, err := db.Exec(sqlQuery, feedback.Comment, feedback.Ip, feedback.Mac)
	if err != nil {
		return "0", err
	}

	return "1", nil
}
