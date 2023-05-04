package models

import (
	"database/sql"
	"fmt"
	. "gamebackend/helpers"
	"log"
	"strings"

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
	Id                 string `json:"Id"`
	Name               string `json:"Name"`
	MiddleName         string `json:"MiddleName"`
	LastName           string `json:"LastName"`
	Email              string `json:"Email"`
	PhoneNumber        string `json:"PhoneNumber"`
	Password           string `json:"Password"`
	BirthDate          string `json:"BirthDate"`
	Gender             string `json:"Gender"`
	NationalId         string `json:"NationalId"`
	PassportId         string `json:"PassportId"`
	Nationality        string `json:"Nationality"`
	VerifyCode         string `json:"VerifyCode"`
	IsVerify           bool   `json:"IsVerify"`
	IsIdentityApproved bool   `json:"IsIdentityApproved"`
	IsEmailApproved    bool   `json:"IsEmailApproved"`
	IsTfaActive        bool   `json:"IsTfaActive"`
	IsPatient          bool   `json:"IsPatient"`
	IsDeleted          bool   `json:"IsDeleted"`
	PhotoUrl           string `json:"PhotoUrl"`
	IsDoctor           bool   `json:"IsDoctor"`
	IsMainUser         bool   `json:"IsMainUser"`
	CreatedAt          string `json:"CreatedAt"`
	UpdatedAt          string `json:"UpdatedAt"`
}

func InsertUser(user User) bool {
	sqlQuery := `
        INSERT INTO user (
            Id, Name, MiddleName, LastName, Email, PhoneNumber,
            Password, BirthDate, Gender, NationalId, PassportId, Nationality, VerifyCode, IsVerify,
            IsIdentityApproved, IsEmailApproved, IsTfaActive, IsPatient, IsDeleted,
            PhotoUrl, IsDoctor, IsMainUser, CreatedAt, UpdatedAt
        ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
    `

	result, err := db.Exec(sqlQuery,
		user.Id, user.Name, user.MiddleName, user.LastName, user.Email, user.PhoneNumber,
		user.Password, user.BirthDate, user.Gender, user.NationalId, user.PassportId, user.Nationality, user.VerifyCode, user.IsVerify,
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

func CheckPhoneNumber(phoneNumber string) (bool, string) {
	var id string
	var count int
	err := db.QueryRow("SELECT user.Id, COUNT(*) FROM user WHERE PhoneNumber LIKE CONCAT('%', ?, '%') GROUP BY user.Id", phoneNumber).Scan(&id, &count)
	if err != nil {
		CheckError(err)
	}
	if count > 0 {
		return true, id
	} else {
		return false, "0"
	}
}

func GetUserById(id string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT * FROM user WHERE Id=?", id).Scan(&user.Id, &user.Name, &user.MiddleName, &user.LastName, &user.Email, &user.PhoneNumber, &user.Password, &user.BirthDate, &user.Gender, &user.NationalId, &user.PassportId, &user.Nationality, &user.VerifyCode, &user.IsVerify, &user.IsIdentityApproved, &user.IsEmailApproved, &user.IsTfaActive, &user.IsPatient, &user.IsDeleted, &user.PhotoUrl, &user.IsDoctor, &user.IsMainUser, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // kullanıcı yok
		}
		fmt.Println(err)
		return nil, err // diğer hatalar
	}

	if user.Id == "" {
		return nil, nil // kullanıcı yok
	}

	return &user, nil // kullanıcı var

}

func CheckVerifyCode(phoneNumber, verifyCode string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user WHERE PhoneNumber AND VerifyCode", phoneNumber, verifyCode).Scan(&count)
	if err != nil {
		CheckError(err)
	}
	if count > 0 {
		return false
	} else {
		return true
	}

}

func UpdateUserFromPhone(PhoneNumber string, updateFields map[string]interface{}) bool {
	sqlQuery := "UPDATE user SET "
	values := make([]interface{}, 0)

	for key, val := range updateFields {
		sqlQuery += key + "=?,"
		values = append(values, val)
	}
	sqlQuery = strings.TrimSuffix(sqlQuery, ",") // remove the last comma
	sqlQuery += " WHERE PhoneNumber = ?"
	values = append(values, PhoneNumber)

	result, err := db.Exec(sqlQuery, values...)
	if err != nil {
		CheckError(err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func UpdateUserFromId(Id string, updateFields map[string]interface{}) bool {
	sqlQuery := "UPDATE user SET "
	values := make([]interface{}, 0)

	for key, val := range updateFields {
		sqlQuery += key + "=?,"
		values = append(values, val)
	}
	sqlQuery = strings.TrimSuffix(sqlQuery, ",") // remove the last comma
	sqlQuery += " WHERE Id = ?"
	values = append(values, Id)

	result, err := db.Exec(sqlQuery, values...)
	if err != nil {
		CheckError(err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected > 0 {
		return true
	} else {
		return false
	}
}
