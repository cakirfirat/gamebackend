package models

import (
	"fmt"
	"log"
)

var myBucket = "taxtrackingbucket"
var accessKey = "AKIATFDGZ4RATBZE2MXH"
var accessSecret = "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"

type Attachment struct {
	Id            string `json:"id"`
	ApplicationId string `json:"application_id"`
	FileName      string `json:"file_name"`
	FileUrl       string `json:"file_url"`
}

func InsertAttachment(attachment Attachment) {
	result, err := db.Exec("INSERT INTO attachment(application_id, file_name, file_url) VALUES(?,?,?)", attachment.ApplicationId, attachment.FileName, attachment.FileUrl)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()

	fmt.Printf("Etkilenen kayıt sayısı(%d)", rowsAffected)
}
func GetAttachment(applicationID string) ([]Attachment, error) {
	rows, err := db.Query("SELECT * FROM attachment WHERE application_id = " + applicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var attachments []Attachment
	for rows.Next() {
		var attachment Attachment
		err := rows.Scan(&attachment.Id, &attachment.ApplicationId, &attachment.FileName, &attachment.FileUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		attachments = append(attachments, attachment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %v", err)
	}

	return attachments, nil
}
