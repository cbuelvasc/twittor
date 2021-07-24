package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cbuelvasc/twittor/database"
	"github.com/cbuelvasc/twittor/jwt"
	"github.com/cbuelvasc/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "appication/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid username or password "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	document, isValid := database.Login(t.Email, t.Password)
	if isValid == false {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWTToken(document)
	if err != nil {
		http.Error(w, "An error occurred while generating the token "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "appication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	})
}
