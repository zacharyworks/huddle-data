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

	return respond(w, *session)
}

// GetSession fetches a session
func GetSessionByState(w http.ResponseWriter, r *http.Request) *shared.AppError {
	sessionID := mux.Vars(r)["id"]

	// Attempt to select to-do from database
	session, e := sql.SelectSessionByState(sessionID)
	if e != nil {
		return e
	}

	return respond(w, *session)
}

// PostSession creates a new session
func PostSession(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return shared.ErrorProcessingBody(err)
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

	return respond(w, *newSession)
}

// PutSession puts a to-do
func PutSession(w http.ResponseWriter, r *http.Request) *shared.AppError {



	var session types.Session
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return shared.ErrorProcessingBody(err)
	}
	if err := json.Unmarshal([]byte(body), &session); err != nil {
		return shared.ErrorProcessingJSON(err)
	}

	// Attempt update
	return sql.UpdateSession(session)
}
