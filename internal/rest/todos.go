package rest

import (
	"../shared"
	"../sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
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

	if e = respond(w, todos); e != nil {
		return e
	}

	return nil
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

	if e = respond(w, todo); e != nil {
		return e
	}
	return nil
}

// PutTodo puts a to-do
func PutTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	todoID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	// Get map of to-do values from body from their keys
	var todo types.Todo
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(body), &todo)

	todo.TodoID = todoID

	// Attempt update
	if e = sql.UpdateTodo(todo); err != nil {
		return e
	}
	return nil
}

// PostTodo posts a to-do
func PostTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		shared.ErrorProcessingBody(err)
	}

	// Get map of to-do values from body from their keys
	var todo types.Todo
	if err := json.Unmarshal([]byte(body), &todo); err != nil {
		return shared.ErrorProcessingJSON(err)
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

	// Form response
	if e = respond(w, newTodo); e != nil {
		return e
	}
	return nil
}

// DeleteTodo deletes a to-do
func DeleteTodo(w http.ResponseWriter, r *http.Request) *shared.AppError {
	todoID, e := getVarAsInt(r, "id")
	if e != nil {
		return e
	}

	// Attempt delete
	if e = sql.RemoveTodo(types.Todo{TodoID: todoID}); e != nil {
		return e
	}
	return nil
}
