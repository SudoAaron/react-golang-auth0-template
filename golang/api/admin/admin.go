package admin

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	responseData := map[string]string{"message": "Welcome to the admin page"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
