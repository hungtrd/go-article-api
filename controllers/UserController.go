package controllers

import (
	"encoding/json"
	"go-article/models"
	"log"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User

	// Parse data
	err := json.NewDecoder(r.Body).Decode(&user)
	log.Panicln(err)

	models.Model.CreateUser(&user)

	json.NewEncoder(w).Encode(user)
}
