package rest

import (
	"github.com/zacharyworks/huddle-data/shared"
	"github.com/zacharyworks/huddle-data/sql"
	"github.com/gorilla/mux"
	types "github.com/zacharyworks/huddle-shared/data"
	"net/http"
)

// AddUserHandlers adds the handlers for user functions
func AddBoardsHandlers(router *mux.Router) {
	router.Handle("/user/{id}/boards", Handler(GetBoardsForUser)).Methods("GET")
	router.Handle("/board/code", Handler(PostBoardJoinCode)).Methods("POST")
	router.Handle("/board/join", Handler(PostJoinBoard)).Methods("POST")
	router.Handle("/board", Handler(PostBoard)).Methods(http.MethodPost)
	router.Handle("/board/member", Handler(DeleteBoardMember)).Methods(http.MethodDelete)
}

func DeleteBoardMember(w http.ResponseWriter, r *http.Request) *shared.AppError {
	var boardMember types.BoardMember
	if e := readBodyIntoType(r.Body, &boardMember); e != nil {
		return e
	}
	return sql.DeleteBoardMember(boardMember)
}

// GetSession fetches a session
func GetBoardsForUser(w http.ResponseWriter, r *http.Request) *shared.AppError {
	userID := mux.Vars(r)["id"]
	boards, e := sql.SelectUsersBoards(userID)
	if e != nil {
		return e
	}
	return respond(w, boards)
}

func PostBoardJoinCode(w http.ResponseWriter, r *http.Request) *shared.AppError {
	var boardJoinCode types.BoardJoinCode
	if e := readBodyIntoType(r.Body, &boardJoinCode); e != nil {
		return e
	}
	if e := sql.InsertBoardJoinCode(boardJoinCode); e != nil {
		return e
	}
	return nil
}

func PostJoinBoard(w http.ResponseWriter, r *http.Request) *shared.AppError {
	var boardJoin types.BoardJoin
	if e := readBodyIntoType(r.Body, &boardJoin); e != nil {
		return e
	}
	board, e := sql.JoinBoard(boardJoin);
	if e != nil {
		return e
	}
	respond(w, board)
	return nil
}

func PostBoard(w http.ResponseWriter, r *http.Request) *shared.AppError {
	var newBoard types.NewBoard
	if e := readBodyIntoType(r.Body, &newBoard); e != nil {
		return e
	}

	board, e := sql.InsertBoard(newBoard.Board)
	if e != nil {
		return e
	}
	if e := sql.InsertBoardMember(
		types.BoardMember{
			BoardFK: board.BoardID,
			UserFK: newBoard.UserFK},
		); e != nil {
		return e
	}


	return respond(w, board)
}