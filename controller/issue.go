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
func (c *Controller) GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := model.GetAllIssues(c.DB)

	if err != nil {
		log.Println("[IssueController@GetIssues] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}

//FindIssue find a issue in database
func (c *Controller) FindIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	issues, err := model.FindIssue(c.DB, id)

	if err != nil {
		log.Println("[IssueController@GetIssues] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}
