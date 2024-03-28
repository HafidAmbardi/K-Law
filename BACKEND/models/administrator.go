package models

import (
	"OTI-inbound/utils/token"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Administrator struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func VerifyPass(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheckAdmin(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Administrator{}

	err = db.Model(Administrator{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Administrator) SaveAdmin(db *gorm.DB) (*Administrator, error) {

	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Administrator{}, errPassword
	}
	u.Password = string(hashedPassword)

	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &Administrator{}, err
	}
	return u, nil
}
