package model

import (
	"log"

	"github.com/alustau/issuer/database"
)

//Issue issue model
type Issue struct {
	ID      string `json:"id,omitempty" db:"id"`
	Number  int    `json:"number,omitempty" db:"number"`
	Details string `json:"details,omitempty" db:"details"`
}

//GetAll add a issue in database
func (issue *Issue) GetAll() (issues []Issue, err error) {
	db, err := database.GetConnection()
	if err != nil {
		log.Println("[issue.GetAll] Error connection: ", err.Error())
		return
	}

	db.Select(&issues, "select * from issue")

	if err != nil {
		log.Println("[issue.GetAll] Error connection: ", err.Error())
		return
	}

	return
}
