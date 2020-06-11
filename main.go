package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-article/controllers"
	"go-article/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests(c controllers.Controller) {
	log.Println("Starting development server at http://127.0.0.1:5000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Route
	myRouter.HandleFunc("/", homePage)

	// User
	myRouter.HandleFunc("/api/users/login", c.UserLogin).Methods("POST")
	myRouter.HandleFunc("/api/users", c.CreateUser).Methods("POST")
	myRouter.HandleFunc("/api/users", c.ListUser).Methods("GET")
	// Get Current User
	myRouter.HandleFunc("/api/user", c.GetCurrentUser).Methods("GET")

	// Article
	myRouter.HandleFunc("/api/articles", c.CreateArticle).Methods("POST")
	myRouter.HandleFunc("/api/articles", c.GetListArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", myRouter))
}

func main() {
	db, dbErr := models.ConnectDB()
	if dbErr != nil {
		log.Println("Connect DB error!")
	}

	contr := controllers.Controller{
		DB: db,
	}
	handleRequests(contr)
}
