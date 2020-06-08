package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	// "io/ioutil"
)

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

type Model struct {
	DB *gorm.DB
}

func (User) TableName() string {
	return "user_models"
}

func (m Model) CreateUser(user User) {
	m.DB.Create(&user)
	log.Print(user)
	fmt.Println("Endpoint Hit: Creating New User")
}
