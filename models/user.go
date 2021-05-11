package models

import (
	"html"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

func (db *DB) CreateUser(user *User) {
	// Trim & Santize string
	user.Username = Santize(user.Username)
	user.Email = Santize(user.Email)
	user.Bio = Santize(user.Bio)
	user.PasswordHash, _ = Hash(Santize(user.PasswordHash))

	// Save User
	db.Create(&user)

	log.Print(user)
}

func (db *DB) FindUserByUsername(username string) (User, error) {
	var user User

	err := db.Where("username = ?", username).First(&user).Error

	return user, err
}

func (db *DB) FindUserByEmail(email string) (User, error) {
	var user User

	err := db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (db *DB) FindUserById(id uint) (User, error) {
	var user User

	err := db.Table("user_models").Where("id = ?", id).First(&user).Error

	return user, err
}

func (db *DB) CheckUserExist(username string) bool {
	var users []User
	db.Table("user_models").Where("username = ? ", username).Find(&users)

	if len(users) > 0 {
		return true
	} else {
		return false
	}
}

func (db *DB) ListUser() *[]User {
	var user []User
	db.Table("user_models").Find(&user)

	return &user
}
