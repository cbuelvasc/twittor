package routers

import (
	"errors"
	"strings"

	"github.com/cbuelvasc/twittor/database"
	"github.com/cbuelvasc/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var UserID string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	pass := []byte("$%&/()=?¿#P422w0rd")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid format token")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return pass, nil
	})
	if err == nil {
		_, found, _ := database.ValidateUser(claims.Email)
		if found == true {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	return claims, false, string(""), err

}
