package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacbriggssagecom/huddle/server/data/internal/sql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"io/ioutil"
	"log"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddUserHandlers(router *mux.Router) {
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user", PostUser).Methods("POST")
}

// GetUser get single todo
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
	id, err := sql.InsertUser(user)
	if err != nil {
		log.Fatal(err)
		return
	}

	// From the returned ID, send back the new todo in JSON
	newUser, err := sql.SelectUserByID(id)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	userJSON, err := json.Marshal(*newUser)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(userJSON)
}
