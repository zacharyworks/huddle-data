package rest

import (
	"../shared"
	"../sql"
	"github.com/gorilla/mux"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddBoardsHandlers(router *mux.Router) {
	router.Handle("/user/{id}/boards", Handler(GetBoardsForUser)).Methods("GET")

}

// GetSession fetches a session
func GetBoardsForUser(w http.ResponseWriter, r *http.Request) *shared.AppError {
	userID := mux.Vars(r)["id"]

	// Attempt to select todo from database
	boards, e := sql.SelectUsersBoards(userID)
	if e != nil {
		return e
	}

	// Convert to JSON
	if e = respond(w, boards); e != nil {
		return e
	}

	return nil
}
