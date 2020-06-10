package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go-article/controllers"
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

	// User
	myRouter.HandleFunc("/api/users/login", controllers.UserLogin).Methods("POST")
	myRouter.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	myRouter.HandleFunc("/api/users", controllers.ListUser).Methods("GET")
	// Get Current User
	myRouter.HandleFunc("/api/user", controllers.GetCurrentUser).Methods("GET")

	// Article
	myRouter.HandleFunc("/api/articles", controllers.CreateArticle).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	handleRequests()
}
