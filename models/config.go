package models

import (
	"database/sql"
	"fmt"
	"log"
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
