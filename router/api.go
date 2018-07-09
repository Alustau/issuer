package router

import (
	"github.com/alustau/issuer/controller"
	"github.com/gorilla/mux"
)

//Router api routes
func Router(c *controller.Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/issue", c.GetIssues).Methods("GET")
	router.HandleFunc("/issue/{id}", c.FindIssue).Methods("GET")
	router.HandleFunc("/issues", c.SaveIssue).Methods("POST")

	return router
}
