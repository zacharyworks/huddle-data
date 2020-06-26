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
func AddSessionHandlers(router *mux.Router) {
	router.HandleFunc("/session/state/{id}", GetSessionByState).Methods("GET")
	router.HandleFunc("/session/id/{id}", GetSessionByID).Methods("GET")
	router.HandleFunc("/session/adduser", PutSession).Methods("PUT")
	router.HandleFunc("/session", PostSession).Methods("POST")
}

// GetSession fetches a session
func GetSessionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get todo ID from url vars
	sessionID := vars["id"]

	// Attempt to select todo from database
	session, err := sql.SelectSessionByID(sessionID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	sessionJSON, err := json.Marshal(*session)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(sessionJSON)
}

// GetSession fetches a session
func GetSessionByState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get todo ID from url vars
	sessionID := vars["id"]

	// Attempt to select todo from database
	session, err := sql.SelectSessionByState(sessionID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	sessionJSON, err := json.Marshal(*session)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(sessionJSON)
}

// PostSession creates a new session
func PostSession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Get map of todo values from body from their keys
	var session types.Session
	json.Unmarshal([]byte(body), &session)

	// Attempt to create todo
	err = sql.InsertSession(session)
	if err != nil {
		println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
	}

	// From the returned ID, send back the new todo in JSON
	newSession, err := sql.SelectSessionByID(session.SessionID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	sessionJSON, err := json.Marshal(*newSession)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(sessionJSON)
}

// PutSession puts a todo
func PutSession(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	// not using vars because they would expose session ID in URL
	var session types.Session
	json.Unmarshal([]byte(body), &session)

	// Attempt update
	err = sql.UpdateSession(session)
	if err != nil {
		log.Fatal(err)
	}
}
