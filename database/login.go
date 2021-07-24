package database

import (
	"github.com/cbuelvasc/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	u, found, _ := ValidateUser(email)
	if found == false {
		return u, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return u, false
	}
	return u, true
}
