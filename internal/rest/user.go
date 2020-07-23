package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"../sql"
	"github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
	"log"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddUserHandlers(router *mux.Router) {
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", PostUser).Methods("POST")
}

// GetUser get single to-do
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get todo ID from url vars
	oauthID := vars["id"]

	// Attempt to select todo from database
	user, err := sql.SelectUser(oauthID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	userJSON, err := json.Marshal(*user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(userJSON)
}

// PostUser creates a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Get map of todo values from body from their keys
	var user types.User
	json.Unmarshal([]byte(body), &user)

	// Attempt to create todo
	newUser, err := sql.InsertUser(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Convert to JSON
	userJSON, err := json.Marshal(*newUser)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(userJSON)
}
