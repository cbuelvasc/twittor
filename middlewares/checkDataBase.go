package middlewares

import (
	"net/http"

	"github.com/cbuelvasc/twittor/database"
)

func CheckDataBase(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(rw, "Lost conecction with database", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
