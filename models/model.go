package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	*gorm.DB
}

func ConnectDB() (*DB, error) {

	db, err := gorm.Open("mysql", "root:hungtran97@tcp(127.0.0.1:3307)/go_article?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed")
	} else {
		log.Println("Connection Established")
	}

	return &DB{db}, err
}
