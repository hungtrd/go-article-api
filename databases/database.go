package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type DB struct {
	*gorm.DB
}

func ConnectDB() *gorm.DB {
	var db DB
	var err error

	db.DB, err = gorm.Open("mysql", "root:hungtran97@tcp(127.0.0.1:3307)/go_article?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection Failed")
	} else {
		log.Println("Connection Established")
	}

	return db.DB
}
