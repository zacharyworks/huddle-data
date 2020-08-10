package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-data/shared"
	"github.com/zacharyworks/huddle-data/sql"
	"github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddUserHandlers(router *mux.Router) {
	router.Handle("/user/{id}", Handler(GetUser)).Methods("GET")
	router.Handle("/user", Handler(PutUser)).Methods("PUT")
	router.Handle("/user", Handler(PostUser)).Methods("POST")

}

// GetUser get single to-do
func GetUser(w http.ResponseWriter, r *http.Request) *shared.AppError {
	oauthID := mux.Vars(r)["id"]

	// Attempt to select to-do from database
	user, e := sql.SelectUser(oauthID)
	if e != nil {
		return e
	}

	if e = respond(w, user); e != nil {
		return e
	}

	return nil
}

// PostUser creates a new user
func PostUser(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return shared.ErrorProcessingBody(err)
	}

	// Get map of user values from body from their keys
	var user types.User
	if err = json.Unmarshal([]byte(body), &user); err != nil {
		return shared.ErrorProcessingJSON(err)
	}

	// Attempt to create todo
	newUser, e := sql.InsertUser(user)
	if e != nil {
		return e
	}

	if e = respond(w, newUser); e != nil {
		return e
	}
	return nil
}

// PostUser creates a new user
func PutUser(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return shared.ErrorProcessingBody(err)
	}

	// Get map of user values from body from their keys
	var user types.User
	if err = json.Unmarshal([]byte(body), &user); err != nil {
		return shared.ErrorProcessingJSON(err)
	}

	// Attempt to update the user
	updatedUser, e := sql.UpdateUser(user)
	if e != nil {
		return e
	}

	if e = respond(w, updatedUser); e != nil {
		return e
	}
	return nil
}