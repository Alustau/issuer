package router

import (
	"log"
	"net/http"

	"github.com/alustau/issuer/controller"
	"github.com/gorilla/mux"
)

//Routes onde estao as rotas
func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/issue", controller.GetIssues).Methods("GET")
	router.HandleFunc("/issue/{id}", controller.FindIssue).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
