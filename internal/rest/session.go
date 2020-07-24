package rest

import (
	"../shared"
	"../sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
	"log"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddSessionHandlers(router *mux.Router) {
	router.Handle("/session/state/{id}", Handler(GetSessionByState)).Methods("GET")
	router.Handle("/session/id/{id}", Handler(GetSessionByID)).Methods("GET")
	router.Handle("/session/adduser", Handler(PutSession)).Methods("PUT")
	router.Handle("/session", Handler(PostSession)).Methods("POST")
}

// GetSession fetches a session
func GetSessionByID(w http.ResponseWriter, r *http.Request) *shared.AppError {
	sessionID := mux.Vars(r)["id"]

	// Attempt to select to-do from database
	session, e := sql.SelectSessionByID(sessionID)
	if e != nil {
		return e
	}

	// Convert to JSON
	if e = respond(w, *session); e != nil {
		return e
	}
	return nil
}

// GetSession fetches a session
func GetSessionByState(w http.ResponseWriter, r *http.Request) *shared.AppError {
	sessionID := mux.Vars(r)["id"]

	// Attempt to select to-do from database
	session, e := sql.SelectSessionByState(sessionID)
	if e != nil {
		log.Fatal(e)
	}

	// Convert to JSON
	if e = respond(w, *session); e != nil {
		return e
	}

	return nil
}

// PostSession creates a new session
func PostSession(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		shared.ErrorProcessingBody(err)
	}

	var session types.Session
	json.Unmarshal([]byte(body), &session)

	// Attempt to create session
	if e := sql.InsertSession(session); e != nil {
		return e
	}

	// From the returned ID, send back the new session
	newSession, e := sql.SelectSessionByID(session.SessionID)
	if e != nil {
		return e
	}

	// Convert to JSON
	if e = respond(w, *newSession); e != nil {
		return e
	}
	return nil
}

// PutSession puts a to-do
func PutSession(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		shared.ErrorProcessingBody(err)
	}

	var session types.Session
	json.Unmarshal([]byte(body), &session)

	// Attempt update
	if e := sql.UpdateSession(session); err != nil {
		return e
	}
	return nil
}
