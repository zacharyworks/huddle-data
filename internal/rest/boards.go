package rest

import (
	"../sql"
	"encoding/json"
	"github.com/gorilla/mux"
	types "github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
	"log"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddBoardsHandlers(router *mux.Router) {
	router.HandleFunc("/user/{id}/boards", GetBoardsForUser).Methods("GET")
	router.HandleFunc("/boards", CreateBoard).Methods("POST")
}

// GetSession fetches a session
func GetBoardsForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get todo ID from url vars
	userID := vars["id"]

	// Attempt to select todo from database
	boards, err := sql.SelectUsersBoards(userID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	boardsJSON, err := json.Marshal(boards)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(boardsJSON)
}

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Get map of todo values from body from their keys
	var board types.Board
	json.Unmarshal([]byte(body), &board)

	// Attempt to create todo
	newBoard, err := sql.InsertBoard(board)
	if err != nil {
		println(err.Error)
		w.WriteHeader(http.StatusBadRequest)
	}


	// Convert to JSON
	newBoardJSON, err := json.Marshal(*newBoard)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(newBoardJSON)
}