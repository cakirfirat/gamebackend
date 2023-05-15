package models

import (
	. "gamebackend/helpers"
)

type Assignment struct {
	DoctorId  string `json:"DoctorId"`
	UserId    string `json:"UserId"`
	IsDeleted string `json:"IsDeleted"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

func SetPatientAssignment(assignment Assignment) bool {

	sqlQuery := `
        INSERT INTO assignment (
            DoctorId, UserId, IsDeleted, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?)
    `

	_, err := db.Exec(sqlQuery, assignment.DoctorId, assignment.UserId, assignment.IsDeleted, assignment.CreatedAt, assignment.UpdatedAt)
	if err != nil {
		CheckError(err)
		return false
	}
	return true
}

func GetPatientAssignments() ([]Assignment, error) {
	var assignments []Assignment

	rows, err := db.Query("SELECT * FROM assignment")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var assignment Assignment
		err := rows.Scan(&assignment.DoctorId, &assignment.UserId, &assignment.IsDeleted, &assignment.CreatedAt, &assignment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}
