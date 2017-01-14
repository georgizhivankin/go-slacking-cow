package main

import (
	"fmt"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func handleRequests() {
	// Initialize a new Mux router
	cowRouter := mux.NewRouter().StrictSlash(true)
	// Register routes
	cowRouter.HandleFunc("/", homePage)
	// Start server and listen for requests
	log.Fatal(http.ListenAndServe(":8010", cowRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	// Load config variables
	configVariables := config.NewConfigVariables()
	fmt.Println(configVariables.ApiKey)
	fmt.Println("Rest API v1.0 - Mux Routers")
	handleRequests()
}
