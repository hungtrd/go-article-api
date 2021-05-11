package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-article/models"
	"go-article/util"
)

func (c Controller) CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Creating New Article")

	var article models.ArticleRequestParam
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}

	if article.Body == "" || article.Description == "" || article.Title == "" {
		err := util.ResponseError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Title, Description and Body can't be blank!",
		}
		util.SendResponseError(w, err)
		return
	}

	// username, errJwt := util.CheckJwt(r)
	username := "hungtran"

	// if errJwt != nil {
	// 	err := util.ResponseError{
	// 		StatusCode: 401,
	// 		Message:    "Unauthorized!",
	// 	}
	// 	util.SendResponseError(w, err)
	// 	return
	// } else {
	user, err := c.DB.FindUserByUsername(username)

	if err == nil {
		userID := user.ID
		arti, err := c.DB.CreateArticle(article, userID)
		if err == nil {
			util.SendResponseData(w, arti)
		}
	}
	// }

}

func (c Controller) GetListArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Get List Articles")

	// username, errJwt := util.CheckJwt(r)
	username := "hungtran"

	// if errJwt != nil {
	// 	err := util.ResponseError{
	// 		StatusCode: 401,
	// 		Message:    "Unauthorized!",
	// 	}
	// 	util.SendResponseError(w, err)
	// 	return
	// } else {
	user, err := c.DB.FindUserByUsername(username)

	if err == nil {
		userID := user.ID
		arti, err := c.DB.GetListArticle(userID)
		if err == nil {
			util.SendResponseData(w, arti)
		}
	}
	// }
}

func (c Controller) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Get Article")

	// var detailRequest models.ArticleDetailRequest
	// err := json.NewDecoder(r.Body).Decode(&detailRequest)

	// if err != nil {
	// 	err := util.ResponseError{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Invalid Format!",
	// 	}
	// 	util.SendResponseError(w, err)
	// 	log.Println(err)
	// 	return
	// }

	// id := detailRequest.ID
	ids := r.FormValue("id")
	id, _ := strconv.Atoi(ids)

	res, err := c.DB.GetArticle(id)

	if err != nil {
		errRes := util.ResponseError{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found!",
		}
		util.SendResponseError(w, errRes)
	} else {
		util.SendResponseData(w, res)
	}

}
func (c Controller) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Update Article")

	var article models.ArticleUpdateRequestParam
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}
	res, err := c.DB.UpdateArticle(article)
	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}
	util.SendResponseData(w, res)
}

func (c Controller) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: Delete Article")

	var detailRequest models.ArticleDetailRequest
	err := json.NewDecoder(r.Body).Decode(&detailRequest)

	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Format!",
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}

	id := detailRequest.ID

	err = c.DB.DeleteArticle(id)
	if err != nil {
		err := util.ResponseError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
		util.SendResponseError(w, err)
		log.Println(err)
		return
	}
	util.SendResponseData(w, "success")
}
