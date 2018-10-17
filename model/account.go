package model

import (
	"database/sql"
	"github.com/tommycpp/Whisper/sqlconnection"
	"log"
)

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Account) StoreIntoDB(db *sqlconnection.SqlConnection) sql.Result {
	res, err := db.Exec("INSERT INTO user(`id`,username,`password`) VALUES (?,?,?)", a.Id, a.Username, a.Password)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Cannot create user")
	}
	return res
}

func (a *Account) CheckIfValid(db *sqlconnection.SqlConnection) bool {
	res, err := db.Query("SELECT * FROM user WHERE username=? AND password=?", a.Username, a.Password)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error when connect to DB")
	}
	if res.Next() {
		return true
	} else {
		log.Fatal(res)
		return false
	}
}