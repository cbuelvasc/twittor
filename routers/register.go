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
		http.Error(w, "Error in received data "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "A password of at least 6 characters must be specified", http.StatusBadRequest)
		return
	}
	_, found, _ := database.ValidateUser(t.Email)
	if found == true {
		http.Error(w, "There is already a registered user with that email", http.StatusBadRequest)
		return
	}
	_, status, err := database.InsertUser(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to save the user "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Unable to insert user record", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
