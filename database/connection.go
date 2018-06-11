package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	//not used by application
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sqlx.DB
)

//GetConnection return db connection
func GetConnection() (db *sqlx.DB, err error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	err = nil

	if db != nil {
		return
	}

	dbstring := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
	db, err = sqlx.Open("mysql", dbstring)

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
