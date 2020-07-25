package rest

import (
	"../shared"
	"../sql"
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
	println(boardJoin.Code)
	println(boardJoin.UserFK)
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