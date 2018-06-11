package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	//not used by application
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

//GetConnection return db connection
func GetConnection() (db *sqlx.DB, err error) {
	err = nil

	if db != nil {
		return
	}

	db, err = sqlx.Open("mysql", "root:alustau321@tcp(localhost:3306)/issuer?parseTime=true")

	if err != nil {
		log.Println("[GetConnection] Error connection: ", err.Error())
		return
	}

	err = db.Ping()

	if err != nil {
		log.Println("[GetConnection] Error ping connection: ", err.Error())
		return
	}

	return
}
