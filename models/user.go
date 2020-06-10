package models

import (
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"go-article/databases"
)

type User struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

func (User) TableName() string {
	return "user_models"
}

func Hash(pass string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(byte), err
}

func CheckPasswordHash(hashedpass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

func CreateUser(user *User) {
	// Trim & Santize string
	user.Username = Santize(user.Username)
	user.Email = Santize(user.Email)
	user.Bio = Santize(user.Bio)
	user.PasswordHash, _ = Hash(Santize(user.PasswordHash))

	// Save User
	DB := db.ConnectDB()
	DB.Create(&user)

	log.Print(user)
}

func FindUserByUsername(username string) (User, error) {
	var user User

	DB := db.ConnectDB()
	err := DB.Where("username = ?", username).First(&user).Error

	return user, err
}

func CheckUserExist(username string) bool {
	DB := db.ConnectDB()
	var users []User
	DB.Table("user_models").Where("username = ? ", username).Find(&users)

	if len(users) > 0 {
		return true
	} else {
		return false
	}
}

func ListUser() *[]User {
	DB := db.ConnectDB()
	var user []User
	DB.Table("user_models").Find(&user)

	return &user
}
