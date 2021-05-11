package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-article/models"
	"go-article/util"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"Image"`
	Token    string  `json:"token"`
}

func (c Controller) UserLogin(w http.ResponseWriter, r *http.Request) {
	var loginParam LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginParam)

	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		util.SendResponseError(w, err)
		log.Println("Error: ", err)
		return
	}

	if loginParam.Email == "" || loginParam.Password == "" {
		err := util.ResponseError{
			StatusCode: 411,
			Message:    "Email and Password can't be blank!",
		}
		util.SendResponseError(w, err)
		log.Println("Error: ", err)
		return
	}

	user, err := c.CheckLogin(loginParam.Email, loginParam.Password)

	if err == nil {
		token, _ := util.CreateJwt(user)
		loginResponse := LoginResponse{
			Username: user.Username,
			Email:    user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
			Token:    token,
		}
		util.SendResponseData(w, loginResponse)
		return
	} else {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Email or Password invalid!",
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Creating New User")

	var user models.User

	// Parse data
	err := json.NewDecoder(r.Body).Decode(&user)

  log.Println(user)
	// Validate data
	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}

	if user.Username == "" || user.PasswordHash == "" || user.Email == "" {
		err := util.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username, Password and Email can't be blank!",
		}
		util.SendResponseError(w, err)
		return
	}

	if c.DB.CheckUserExist(user.Username) {
		err := util.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Username already exists!",
		}
		util.SendResponseError(w, err)
	} else {
		c.DB.CreateUser(&user)
		err := util.ResponseError{
			StatusCode: http.StatusOK,
			Message:    "Register Successfully!",
		}
		util.SendResponseError(w, err)
	}
}

func (c Controller) ListUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get List User")

	w.Header().Set("Content-Type", "application/json")

	user := c.DB.ListUser()

	json.NewEncoder(w).Encode(user)
}

func (c Controller) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Get Current User")

	w.Header().Set("Contend-Type", "application/json")

	username, errJwt := util.CheckJwt(r)

	if errJwt != nil {
		err := util.ResponseError{
			StatusCode: 401,
			Message:    "Unauthorized requests!",
		}
		util.SendResponseError(w, err)
		return
	}

	user, _ := c.DB.FindUserByUsername(username)

	// Reset Token Exp
	token, _ := util.CreateJwt(user)
	loginResponse := LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
		Token:    token,
	}
	util.SendResponseData(w, loginResponse)
}

func (c Controller) CheckLogin(email, password string) (models.User, error) {
	var user models.User
	c.DB.Where("email = ?", email).First(&user)

	err := models.CheckPasswordHash(user.PasswordHash, password)

	log.Println(user)

	return user, err
}
