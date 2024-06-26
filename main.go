package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create new route
	router := mux.NewRouter()

	//specify endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, welcome to first api endpoint %s\n", r.URL.Path)
	})
	// endpoints, handler functions and HTTP method
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")

	fmt.Println("server is up and running on port 80....")
	log.Fatal(http.ListenAndServe(":80", router))

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}
