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
func (issue *Issue) all(c *database.Connection) (issues []Issue, err error) {
	sql := "select * from issue"

	if err := c.DB.Select(&issues, sql); err != nil {
		log.Println("[issue.all] Error when trying get all issues: ", err.Error())
	}

	return
}

//Find add a issue in database
func (issue *Issue) find(c *database.Connection, id int) (row Issue, err error) {
	sql := "select * from issue where id = ?"

	if err := c.DB.Get(&row, sql, id); err != nil {
		log.Println("[issue.find] Error when trying find a issue: ", err.Error())
	}

	return
}

//GetAllIssues ...
func GetAllIssues(c *database.Connection) ([]Issue, error) {
	issue := &Issue{}

	return issue.all(c)
}

//GetAllIssues ...
func FindIssue(c *database.Connection, id int) (Issue, error) {
	issue := &Issue{}

	return issue.find(c, id)
}
