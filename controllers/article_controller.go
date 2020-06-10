package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-article/models"
	"go-article/ulti"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Creating New Article")

	var article models.ArticleRequestParam
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		err := ulti.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		ulti.SendResponseError(w, err)
		log.Println(err)
		return
	}

	if article.Body == "" || article.Description == "" || article.Title == "" {
		err := ulti.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Title, Description and Body can't be blank!",
		}
		ulti.SendResponseError(w, err)
		return
	}

	username, errJwt := ulti.CheckJwt(r)

	if errJwt != nil {
		user, err := models.FindUserByUsername(username)

		if err != nil {
			user_id := user.ID
			models.CreateArticle(article, user_id)
		}
	}

}
