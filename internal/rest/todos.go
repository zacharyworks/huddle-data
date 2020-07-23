package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"../sql"
	"github.com/zacharyworks/huddle-shared/data"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// AddTodoHandlers adds handlers for to-do functions
func AddTodoHandlers(router *mux.Router) {
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/board/{id}/todos", GetBoardTodo).Methods("GET")
	router.HandleFunc("/todos", PostTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", PutTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
}

// GetTodos get to-dos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Get todos
	todos, err := sql.GetAllTodo()
	if err != nil {
		log.Fatal(err)
	}
	// Convert to JSON
	todosJSON, err := json.Marshal(*todos)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(todosJSON)
}

// GetTodos get to-dos
func GetBoardTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get to-do ID from url vars
	boardID, err := strconv.Atoi(vars["id"])
	// Get todos
	todos, err := sql.GetTodosForBoard(boardID)
	if err != nil {
		log.Fatal(err)
	}
	// Convert to JSON
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(todosJSON)
}

// GetTodo get single to-do
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get to-do ID from url vars
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Attempt to select to-do from database
	todo, err := sql.SelectTodo(todoID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	todoJSON, err := json.Marshal(*todo)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(todoJSON)
}

// PutTodo puts a to-do
func PutTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	// Get to-do ID from url vars
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Get map of to-do values from body from their keys
	var todo types.Todo
	json.Unmarshal([]byte(body), &todo)

	todo.TodoID = todoID

	// Attempt update
	err = sql.UpdateTodo(todo)
	if err != nil {
		log.Fatal(err)
	}
}

// PostTodo posts a to-do
func PostTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Get map of to-do values from body from their keys
	var todo types.Todo
	json.Unmarshal([]byte(body), &todo)

	// Attempt to create to-do
	id, err := sql.InsertTodo(todo)
	if err != nil {
		log.Fatal(err)
		return
	}

	// From the returned ID, send back the new to-do in JSON
	newTodo, err := sql.SelectTodo(id)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to JSON
	todoJSON, err := json.Marshal(*newTodo)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(todoJSON)
}

// DeleteTodo deletes a to-do
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Get todo ID from url vars
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Attempt delete
	err = sql.RemoveTodo(types.Todo{TodoID: todoID})
	if err != nil {
		log.Fatal(err)
	}
}
