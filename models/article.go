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

type ArticleDetailRequest struct {
	ID int `gorm:"id"`
}

type ArticleRequestParam struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type ArticleUpdateRequestParam struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

type ArticleResponseParam struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Slug        string              `json:"slug"`
	Description string              `json:"description"`
	Body        string              `json:"body"`
	Author      AuthorResponseParam `json:"author"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
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

func (db *DB) CreateArticle(article ArticleRequestParam, authorID uint) (ArticleResponseParam, error) {
	var articleModel Article
	// Trim & Santize string
	articleModel.Slug = Santize(article.Slug)
	articleModel.Title = Santize(article.Title)
	articleModel.Description = Santize(article.Description)
	articleModel.Body = Santize(article.Body)
	articleModel.AuthorID = authorID
	articleModel.CreatedAt = time.Now()
	articleModel.UpdatedAt = time.Now()

	log.Println("Article Request: ", article)
	log.Println("Article Model: ", articleModel)

	// Save Article
	err := db.Create(&articleModel).Error
	user, _ := db.FindUserById(authorID)

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
		CreatedAt: articleModel.CreatedAt,
		UpdatedAt: articleModel.UpdatedAt,
	}

	return articleRes, err
}

// func

func (db *DB) GetListArticle(authorID uint) (ListArticleResponseParam, error) {
	var articles []Article
	db.Table("articles").Where("author_id = ?", authorID).Find(&articles)
	user, err := db.FindUserById(authorID)

	var articleRes []ArticleResponseParam
	for _, v := range articles {
		articleRes = append(articleRes, ArticleResponseParam{
			ID:          v.ID,
			Title:       v.Title,
			Slug:        v.Slug,
			Body:        v.Body,
			Description: v.Description,
			Author: AuthorResponseParam{
				Username: user.Username,
				Bio:      user.Bio,
				Image:    *user.Image,
			},
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	listArticle := ListArticleResponseParam{
		Articles:      articleRes,
		ArticlesCount: len(articleRes),
	}

	return listArticle, err
}

func (db *DB) GetArticle(id int) (ArticleResponseParam, error) {
	var article Article
	db.Table("articles").Where("id = ?", id).First(&article)
	log.Println(article)
	user, err := db.FindUserById(article.AuthorID)

	var articleRes ArticleResponseParam
	if err == nil {
		articleRes = ArticleResponseParam{
			ID:          article.ID,
			Title:       article.Title,
			Slug:        article.Slug,
			Body:        article.Body,
			Description: article.Description,
			Author: AuthorResponseParam{
				Username: user.Username,
				Bio:      user.Bio,
				Image:    *user.Image,
			},
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		}

	}
	return articleRes, err
}

func (db *DB) UpdateArticle(article ArticleUpdateRequestParam) (ArticleResponseParam, error) {
	var updatedArticle Article
	errUpdate := db.Table("articles").Where("id = ?", article.ID).First(&updatedArticle).Error

	log.Println(updatedArticle)
	user, err := db.FindUserById(updatedArticle.AuthorID)

	updatedArticle.Title = article.Title
	updatedArticle.Body = article.Body
	updatedArticle.Slug = article.Slug
	updatedArticle.Description = article.Description

	db.Save(&updatedArticle)

	var res ArticleResponseParam
	if err == nil {
		res = ArticleResponseParam{
			ID:          updatedArticle.ID,
			Title:       updatedArticle.Title,
			Slug:        updatedArticle.Slug,
			Body:        updatedArticle.Body,
			Description: updatedArticle.Description,
			Author: AuthorResponseParam{
				Username: user.Username,
				Bio:      user.Bio,
				Image:    *user.Image,
			},
			CreatedAt: updatedArticle.CreatedAt,
			UpdatedAt: updatedArticle.UpdatedAt,
		}
	}

	return res, errUpdate
}

func (db *DB) DeleteArticle(id int) error {
	err := db.Delete(&Article{}, id).Error

	return err
}
