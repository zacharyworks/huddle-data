package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zacharyworks/huddle-data/rest"
	"github.com/zacharyworks/huddle-shared/db"
	"io/ioutil"
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
	db.ConnectDB(readAuthCredentials())
	router := mux.NewRouter().StrictSlash(true)
	addHandlers(router)
	rest.AddTodoHandlers(router)
	rest.AddUserHandlers(router)
	rest.AddSessionHandlers(router)
	rest.AddBoardsHandlers(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func readAuthCredentials() (db.Credentials) {
	file, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	credentials := make(map[string]string)
	err = json.Unmarshal([]byte(file), &credentials)
	if err != nil {
		log.Fatal(err)
	}

	dbCredentials := db.Credentials{
		DbURL: credentials["DbURL"],
		DbPort: credentials["DbPort"],
		DbName: credentials["DbName"],
		DbUser: credentials["DbUser"],
		DbPass: credentials["DbPass"],

	}

	return dbCredentials
}
