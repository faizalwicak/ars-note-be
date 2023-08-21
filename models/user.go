package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"size:100;not null;unique" json:"Email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Books    []Book
}

// Username  string `gorm:"size:255;not null;unique" json:"username"`

func GetUserById(uid uint) (User, error) {
	var user User

	db, err := Setup()
	if err != nil {
		log.Println(err)
		return User{}, err
	}

	if err := db.Where("id=?", uid).Find(&user).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
