package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-article/models"
	"go-article/ulti"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Creating New User")

	var user models.User

	// Parse data
	err := json.NewDecoder(r.Body).Decode(&user)

	// Validate data
	if err != nil {
		err := ulti.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		ulti.SendResponseError(w, err)
		log.Println(err)
		return
	}

	if user.Username == "" || user.PasswordHash == "" {
		err := ulti.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username and Password can't be blank!",
		}
		ulti.SendResponseError(w, err)
		return
	}

	if models.CheckUserExist(user.Username) {
		err := ulti.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username existed!",
		}
		ulti.SendResponseError(w, err)
	} else {
		models.CreateUser(&user)
	}
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Get List User")

	user := models.ListUser()

	json.NewEncoder(w).Encode(user)
}
