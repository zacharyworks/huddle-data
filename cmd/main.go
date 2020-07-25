package main

import (
	"../internal/rest"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-shared/db"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Huddle Database Service")
}

func addHandlers(router *mux.Router) {
	router.HandleFunc("/", homeLink)
}

func main() {
	db.ConnectDB()
	router := mux.NewRouter().StrictSlash(true)
	addHandlers(router)
	rest.AddTodoHandlers(router)
	rest.AddUserHandlers(router)
	rest.AddSessionHandlers(router)
	rest.AddBoardsHandlers(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
