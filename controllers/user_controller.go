package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-article/models"
	"go-article/ulti"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"Image"`
	Token    string  `json:"token"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	var loginParam LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginParam)

	if err != nil {
		err := ulti.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		ulti.SendResponseError(w, err)
		log.Println("Error: ", err)
		return
	}

	if loginParam.Username == "" || loginParam.Password == "" {
		err := ulti.ResponseError{
			StatusCode: 411,
			Message:    "Username and Password can't be blank!",
		}
		ulti.SendResponseError(w, err)
		log.Println("Error: ", err)
		return
	}

	user, err := ulti.CheckLogin(loginParam.Username, loginParam.Password)

	if err == nil {
		token, _ := ulti.CreateJwt(user)
		loginResponse := LoginResponse{
			Username: user.Username,
			Email:    user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
			Token:    token,
		}
		ulti.SendResponseData(w, loginResponse)
		return
	} else {
		err := ulti.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Username or Password invalid!",
		}
		ulti.SendResponseError(w, err)
		log.Println(err)
		return
	}
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

	if user.Username == "" || user.PasswordHash == "" || user.Email == "" {
		err := ulti.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username, Password and Email can't be blank!",
		}
		ulti.SendResponseError(w, err)
		return
	}

	if models.CheckUserExist(user.Username) {
		err := ulti.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username already exists!",
		}
		ulti.SendResponseError(w, err)
	} else {
		models.CreateUser(&user)
		err := ulti.ResponseError{
			StatusCode: http.StatusOK,
			Message:    "Register Successfully!",
		}
		ulti.SendResponseError(w, err)
	}
}

func ListUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Get List User")

	user := models.ListUser()

	json.NewEncoder(w).Encode(user)
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contend-Type", "application/json")
	fmt.Println("Endpoint Hit: Get Current User")

	username, errJwt := ulti.CheckJwt(r)

	if errJwt != nil {
		err := ulti.ResponseError{
			StatusCode: 401,
			Message:    "Unauthorized requests!",
		}
		ulti.SendResponseError(w, err)
		return
	}

	user := models.FindUserByUsername(username)

	// Reset Token Exp
	token, _ := ulti.CreateJwt(user)
	loginResponse := LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
	}
	ulti.SendResponseData(w, loginResponse)
}
