package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	//not used by application
	_ "github.com/go-sql-driver/mysql"
)

var (
	connection *Connection
)

func init() {
	connection = new(Connection)
}

//Connection ...
type Connection struct {
	DB *sqlx.DB
}

//Connect ...
func (mysql *Connection) Connect(url string) error {
	db, err := sqlx.Open("mysql", url)

	if err != nil {
		log.Println("[Connect] Trying to connect database: ", err.Error())
		return err
	}

	err = db.Ping()

	if err != nil {
		log.Println("[Connect] Error ping connection: ", err.Error())
		return err
	}

	mysql.DB = db

	return nil
}

//GetConnection Returns the current global connection object
func connect() *Connection {
	return connection
}

// Close the stored session
func (mysql *Connection) Close() {
	mysql.DB.Close()
	//m.session.Close()
}

//GetConnection return Connection
func GetConnection(host string, port string, user string, password string, dbname string) (*Connection, error) {
	mysql := connect()

	if mysql.DB == nil {
		mysql.Connect(parseURL(host, port, user, password, dbname))
	}

	return mysql, nil
}

//parseURL ...
func parseURL(host string, port string, user string, password string, dbname string) string {
	return user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
}
