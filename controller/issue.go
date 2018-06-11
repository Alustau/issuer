package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alustau/issuer/model"
)

//GetIssue return all issues
func GetIssue(w http.ResponseWriter, r *http.Request) {
	issue := model.Issue{}
	issues, err := issue.GetAll()

	if err != nil {
		log.Println("[IssueController@GetIssue] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}
