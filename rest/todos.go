package rest

import (
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-data/shared"
	"github.com/zacharyworks/huddle-data/sql"
	"github.com/zacharyworks/huddle-shared/data"
	"net/http"
)

// AddTodoHandlers adds handlers for to-do functions to a provided router
func AddTodoHandlers(router *mux.Router) {
	router.Handle("/board/{id}/todos", Handler(GetBoardTodo)).Methods("GET")
	router.Handle("/todos", Handler(PostTodo)).Methods("POST")
	router.Handle("/todos/{id}", Handler(GetTodo)).Methods("GET")
	router.Handle("/todos/{id}", Handler(PutTodo)).Methods("PUT")
	router.Handle("/todos/{id}", Handler(DeleteTodo)).Methods("DELETE")
}


func GetBoardTodo(w http.ResponseWriter, r *http.Request) *shared.AppError  {
	boardID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	todos, e := sql.GetTodosForBoard(boardID)
	if e != nil {
		return e
	}

	return respond(w, todos)
}


// GetTodo get single to-do
func GetTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	todoID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	// Attempt to select to-do from database
	todo, e := sql.SelectTodo(todoID)
	if e != nil {
		return e
	}

	return respond(w, todo)
}

// PutTodo puts a to-do
func PutTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	todoID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	// Get map of to-do values from body from their keys
	var todo types.Todo
	if e := readBodyIntoType(r.Body, &todo); e != nil {
		return e
	}

	todo.TodoID = todoID

	// Attempt update
	return sql.UpdateTodo(todo)
}

// PostTodo posts a to-do
func PostTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	// Get map of to-do values from body from their keys
	var todo types.Todo
	if e := readBodyIntoType(r.Body, &todo); e != nil {
		return e
	}
	// Attempt to create to-do
	id, e := sql.InsertTodo(todo)
	if e != nil {
		return e
	}

	// From the returned ID, send back the new to-do in JSON
	newTodo, e := sql.SelectTodo(id)
	if e != nil {
		return e
	}

	return respond(w, newTodo)
}

// DeleteTodo deletes a to-do
func DeleteTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	todoID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	// Attempt delete
	return sql.RemoveTodo(types.Todo{TodoID: todoID})
}
