package routers

import (
	"encoding/json"
	"net/http"

	"github.com/cbuelvasc/twittor/database"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	perfil, err := database.SearchProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find the record "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
