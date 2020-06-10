package models

import (
	"log"
	"time"

	"go-article/databases"
	"go-article/ulti"
)

type Article struct {
	ID          uint      `gorm:"primary_key"`
	Slug        string    `gorm:"column:slug;unique_index"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description;size:2048"`
	Body        string    `gorm:"column:body;size:2048"`
	AuthorID    uint      `gorm:"column:author_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type ArticleRequestParam struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (Article) TableName() string {
	return "articles"
}

func CreateArticle(article ArticleRequestParam, author_id uint) {
	var articleModel Article
	// Trim & Santize string
	articleModel.Slug = ulti.Santize(article.Slug)
	articleModel.Title = Santize(article.Title)
	articleModel.Description = Santize(article.Description)
	articleModel.Body = Santize(article.Body)
	articleModel.AuthorID = author_id
	articleModel.CreatedAt = time.Now()
	articleModel.UpdatedAt = time.Now()

	// Save User
	DB := db.ConnectDB()
	DB.Create(&articleModel)

	log.Println(articleModel)
}
