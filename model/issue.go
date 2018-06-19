package model

import (
	"encoding/json"
	"log"

	"github.com/alustau/issuer/request"

	"github.com/alustau/issuer/database"
)

//Issue issue model
type Issue struct {
	ID     string `json:"id,omitempty" db:"id"`
	Number int    `json:"number,omitempty" db:"number"`
	Issue  string `json:"issue,omitempty" db:"issue"`
	Action string `json:"action,omitempty" db:"action"`
	Labels string `json:"labels,omitempty" db:"labels"`
}

//GetAll add a issue in database
func (issue *Issue) all(c *database.Connection) (issues []Issue, err error) {
	sql := "select * from issue"

	if err := c.DB.Select(&issues, sql); err != nil {
		log.Println("[model.issue.all] Error when trying get all issues: ", err.Error())
	}

	return
}

func (issue *Issue) find(c *database.Connection, id int) (issues []Issue, err error) {
	sql := "select * from issue where number = ?"

	if err := c.DB.Select(&issues, sql, id); err != nil {
		log.Println("[model.issue.find] Error when trying find a issue: ", err.Error())
	}

	return
}

func (issue *Issue) insert(c *database.Connection) (int64, error) {
	sql := "insert into issue (number, issue, action) values (?, ?, ?)"
	result, err := c.DB.Exec(sql, issue.Number, issue.Issue, issue.Action)

	if err != nil {
		log.Println("[model.Issue.insert] ", err.Error())
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("[model.Issue.insert] ", err.Error())
		return rows, err
	}

	return rows, nil
}

//GetAllIssues ...
func GetAllIssues(c *database.Connection) ([]Issue, error) {
	issue := &Issue{}

	return issue.all(c)
}

//FindIssue ...
func FindIssue(c *database.Connection, id int) ([]Issue, error) {
	issue := &Issue{}

	return issue.find(c, id)
}

//InsertIssue ...
func InsertIssue(c *database.Connection, params *request.IssueRequest) (bool, error) {
	issue := &Issue{}

	bytesIssue, err := json.Marshal(params.Issue)

	if err != nil {
		log.Println("[model.InsertIssue] json.Marshal(params.Issue) ", err.Error())
		return false, err
	}

	bytesLabels, err := json.Marshal(params.Issue.Labels)

	if err != nil {
		log.Println("[model.InsertIssue] json.Marshal(params.Issue.Labels)", err.Error())
		return false, err
	}

	issue.Action = params.Action
	issue.Labels = string(bytesLabels)
	issue.Issue = string(bytesIssue)
	issue.Number = params.Issue.Number

	if _, err := issue.insert(c); err != nil {
		return true, err
	}

	return true, nil
}
