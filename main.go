package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// DB connection

const (
	dbname = "go-crud"
)

var db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-crud")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")
}
func main() {

	// Create new route
	router := mux.NewRouter()

	//specify endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, welcome to first api endpoint %s\n", r.URL.Path)
	})
	// endpoints, handler functions and HTTP method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/ping-db", dbPing).Methods("GET")

	fmt.Println("server is up and running on port 7777....")
	log.Fatal(http.ListenAndServe(":7777", router))

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}

func dbPing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connecte1")

}
