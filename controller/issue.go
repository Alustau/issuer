package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alustau/issuer/model"
	"github.com/gorilla/mux"
)

//GetIssues return all issues
func GetIssues(w http.ResponseWriter, r *http.Request) {
	issue := model.Issue{}
	issues, err := issue.GetAll()

	if err != nil {
		log.Println("[IssueController@GetIssues] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}

//FindIssue return all issues
func FindIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	issue := model.Issue{}
	issues, err := issue.Find(id)

	if err != nil {
		log.Println("[IssueController@GetIssues] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}
