package models

import (
	"log"
	"time"
	// "go-article/databases"
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

type ArticleResponseParam struct {
	Title       string              `json:"title"`
	Slug        string              `json:"slug"`
	Description string              `json:"description"`
	Body        string              `json:"body"`
	Author      AuthorResponseParam `json:"author"`
}

type AuthorResponseParam struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type ListArticleResponseParam struct {
	Articles      []ArticleResponseParam `json:"articles"`
	ArticlesCount int                    `json:"articleCount"`
}

func (Article) TableName() string {
	return "articles"
}

func (db *DB) CreateArticle(article ArticleRequestParam, author_id uint) (ArticleResponseParam, error) {
	var articleModel Article
	// Trim & Santize string
	articleModel.Slug = Santize(article.Slug)
	articleModel.Title = Santize(article.Title)
	articleModel.Description = Santize(article.Description)
	articleModel.Body = Santize(article.Body)
	articleModel.AuthorID = author_id
	articleModel.CreatedAt = time.Now()
	articleModel.UpdatedAt = time.Now()

	log.Println("Article Request: ", article)
	log.Println("Article Model: ", articleModel)

	// Save Article
	err := db.Create(&articleModel).Error
	user, _ := db.FindUserById(author_id)

	log.Println(articleModel)
	articleRes := ArticleResponseParam{
		Title:       articleModel.Title,
		Slug:        articleModel.Slug,
		Body:        articleModel.Body,
		Description: articleModel.Description,
		Author: AuthorResponseParam{
			Username: user.Username,
			Bio:      user.Bio,
			Image:    *user.Image,
		},
	}

	return articleRes, err
}

func (db *DB) GetListArticle(author_id uint) (ListArticleResponseParam, error) {
	var articles []Article
	db.Table("articles").Where("author_id = ?", author_id).Find(&articles)
	user, err := db.FindUserById(author_id)

	var articleRes []ArticleResponseParam
	for _, v := range articles {
		articleRes = append(articleRes, ArticleResponseParam{
			Title:       v.Title,
			Slug:        v.Slug,
			Body:        v.Body,
			Description: v.Description,
			Author: AuthorResponseParam{
				Username: user.Username,
				Bio:      user.Bio,
				Image:    *user.Image,
			},
		})
	}

	listArticle := ListArticleResponseParam{
		Articles:      articleRes,
		ArticlesCount: len(articleRes),
	}

	return listArticle, err
}
