package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var db *gorm.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests() {
	log.Println("Starting development server at http://127.0.0.1:5000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Route
	myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/users", UserModel.CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	db, err = gorm.Open("mysql", "root:hungtran97@tcp(127.0.0.1:3307)/go_article?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed")
	} else {
		log.Println("Connection Established")
	}

	handleRequests()
}
