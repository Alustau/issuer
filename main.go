package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alustau/issuer/controller"
	"github.com/alustau/issuer/database"
	"github.com/alustau/issuer/router"
	"github.com/joho/godotenv"
)

const (
	port         = 8000
	readTimeOut  = 10 * time.Second
	writeTimeout = 10 * time.Second
)

var (
	dbHost     string
	dbPort     string
	dbname     string
	dbUsername string
	dbPassword string
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.StringVar(&dbHost, "localhost", os.Getenv("DB_HOST"), "database host")
	flag.StringVar(&dbPort, "3306", os.Getenv("DB_PORT"), "database port")
	flag.StringVar(&dbname, "issuer", os.Getenv("DB_DATABASE"), "database name")
	flag.StringVar(&dbUsername, "root", os.Getenv("DB_USER"), "database user")
	flag.StringVar(&dbPassword, "", os.Getenv("DB_PASSWORD"), "database password")
}

func main() {
	flag.Parse()

	fmt.Printf("===========================================================\n")
	fmt.Printf("==================== Github Issuer=========================\n")
	fmt.Printf("===========================================================\n")

	log.Println("Starting Connection Database...")

	mysql, err := database.GetConnection(dbHost, dbPort, dbUsername, dbPassword, dbname)

	if err != nil {
		log.Println("Error in MySql Connection. Error:", err)
		os.Exit(1)
	}

	defer mysql.Close()

	c := &controller.Controller{DB: mysql}
	routes := router.Router(c)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routes,
		ReadTimeout:  readTimeOut,
		WriteTimeout: writeTimeout,
	}

	log.Println("Starting Server in ", port)

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error in ListenAndServe. Error:", err)
		os.Exit(1)
	}

	fmt.Println("Server has started")
}
