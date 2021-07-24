package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cbuelvasc/twittor/database"
	"github.com/cbuelvasc/twittor/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "A password of at least 6 characters must be specified", 400)
		return
	}
	_, finded, _ := database.ValidateUser(t.Email)
	if finded == true {
		http.Error(w, "There is already a registered user with that email", 400)
		return
	}
	_, status, err := database.InsertUser(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to save the user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Unable to insert user record", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
