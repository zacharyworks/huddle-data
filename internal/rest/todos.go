package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/zacbriggssagecom/huddle/server/data/internal/sql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// AddTodoHandlers adds handlers for todo functions
func AddTodoHandlers(router *mux.Router) {
	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", PostTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", PutTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")
}

// GetTodos get todos
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

// GetTodo get single todo
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// Get todo ID from url vars
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Attempt to select todo from database
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

// PutTodo puts a todo
func PutTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	// Get todo ID from url vars
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Get map of todo values from body from their keys
	var todo types.Todo
	json.Unmarshal([]byte(body), &todo)

	todo.TodoID = todoID

	// Attempt update
	err = sql.UpdateTodo(todo)
	if err != nil {
		log.Fatal(err)
	}
}

// PostTodo posts a todo
func PostTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	// Get map of todo values from body from their keys
	var todo types.Todo
	json.Unmarshal([]byte(body), &todo)

	// Attempt to create todo
	id, err := sql.InsertTodo(todo)
	if err != nil {
		log.Fatal(err)
		return
	}

	// From the returned ID, send back the new todo in JSON
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

// DeleteTodo deletes a todo
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
