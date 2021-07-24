package jwt

import (
	"time"

	"github.com/cbuelvasc/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(t models.User) (string, error) {
	pass := []byte("$%&/()=?¿#P422w0rd")
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.Lastname,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(pass)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
