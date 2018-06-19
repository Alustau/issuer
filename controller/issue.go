package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alustau/issuer/model"
	"github.com/alustau/issuer/request"
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
		log.Println("[IssueController@FindIssue] ", err.Error())
	}

	json.NewEncoder(w).Encode(issues)
}

//SaveIssue find a issue in database
func (c *Controller) SaveIssue(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("[IssueController@SaveIssue] ParseForm: ", err.Error())
		return
	}

	params := new(request.IssueRequest)

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		log.Println("[IssueController@SaveIssue] Decoder: ", err.Error())
		return
	}

	if _, err := model.InsertIssue(c.DB, params); err != nil {
		log.Println("[IssueController@SaveIssue]", err)
	}

	json.NewEncoder(w).Encode(params)
}
