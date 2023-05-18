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

func GetPatientAssignments(userId string) ([]Assignment, error) {
	var assignments []Assignment

	rows, err := db.Query("SELECT * FROM assignment WHERE DoctorId=?", userId)
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

func GetUserByAssignment(id string) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT * FROM user WHERE Id=?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.MiddleName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Password, &user.BirthDate, &user.Gender, &user.NationalId, &user.PassportId, &user.Nationality, &user.VerifyCode, &user.IsVerify, &user.IsIdentityApproved, &user.IsEmailApproved, &user.IsTfaActive, &user.IsPatient, &user.IsDeleted, &user.PhotoUrl, &user.IsDoctor, &user.IsMainUser, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
