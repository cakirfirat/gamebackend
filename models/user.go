package models

import (
	"database/sql"
	"fmt"
	. "gamebackend/helpers"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost = "157.230.56.58"
	dbPort = 3306
	dbUser = "user"
	dbPass = "Emc_1486374269_Emc"
	dbName = "gamebackend"
)

var db *sql.DB

func init() {
	var err error
	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbConnString)
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name               string    `json:"Name"`
	MiddleName         string    `json:"MiddleName"`
	LastName           string    `json:"LastName"`
	Email              string    `json:"Email"`
	PhoneNumber        string    `json:"PhoneNumber"`
	Password           string    `json:"Password"`
	BirthDate          string    `json:"BirthDate"`
	Gender             string    `json:"Gender"`
	NationalId         string    `json:"NationalId"`
	PassportId         string    `json:"PassportId"`
	Nationality        string    `json:"Nationality"`
	IsIdentityApproved bool      `json:"IsIdentityApproved"`
	IsEmailApproved    bool      `json:"IsEmailApproved"`
	IsTfaActive        bool      `json:"IsTfaActive"`
	IsPatient          bool      `json:"IsPatient"`
	IsDeleted          bool      `json:"IsDeleted"`
	PhotoUrl           string    `json:"PhotoUrl"`
	IsDoctor           bool      `json:"IsDoctor"`
	IsMainUser         bool      `json:"IsMainUser"`
	CreatedAt          time.Time `json:"CreatedAt"`
	UpdatedAt          time.Time `json:"UpdatedAt"`
}

func InsertUser(user User) bool {
	sqlQuery := `
        INSERT INTO users (
            Name, MiddleName, LastName, Email, PhoneNumber, 
            Password, BirthDate, NationalId, PassportId, Nationality, 
            IsIdentityApproved, IsEmailApproved, IsTfaActive, IsPatient, IsDeleted,
			PhotoUrl, IsDoctor, IsMainUser, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,)
    `

	result, err := db.Exec(sqlQuery,
		user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber,
		user.Password, user.BirthDate, user.NationalId, user.PassportId, user.Nationality,
		user.IsIdentityApproved, user.IsEmailApproved, user.IsTfaActive, user.IsPatient, user.IsDeleted,
		user.PhotoUrl, user.IsDoctor, user.IsMainUser, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		CheckError(err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected > 0 {
		return true
	} else {
		return false
	}

}
