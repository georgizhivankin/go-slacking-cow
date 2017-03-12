package main

import (
	"fmt"
	"github.com/georgizhivankin/go-slacking-cow/config"
	"github.com/georgizhivankin/go-slacking-cow/helpers"
	"github.com/georgizhivankin/go-slacking-cow/quotes"
	"github.com/georgizhivankin/go-slacking-cow/repositories"
	"github.com/gorilla/mux"
	"html"
	"net/http"
)

func handleRequests() {
	// Initialize a new Mux router
	cowRouter := mux.NewRouter().StrictSlash(true)
	// Register routes
	cowRouter.HandleFunc("/", homePage)
	// Start server and listen for requests
	helpers.CheckAndLogError(http.ListenAndServe(":8010", cowRouter), "fatal")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	// Load config variables
	config, err := config.LoadEnv()
	helpers.CheckAndLogError(err, "panic")
	fmt.Printf("Welcome to %s %s\n", config.AppName, config.AppVersion)
	quote := quotes.GetRandomQuote()

	savedQuote := repositories.SaveQuote(quote)
	if savedQuote == true {
		fmt.Println("Quote successfully saved into the database.")
	}
	handleRequests()
}
