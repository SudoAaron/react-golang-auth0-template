package private

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	// Get the token claims from the request context
	token := r.Context().Value(jwtmiddleware.ContextKey{})
	if token == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Unauthorized."}`))
		return
	}

	claims, ok := token.(*validator.ValidatedClaims)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Unauthorized."}`))
		return
	}

	sub := claims.RegisteredClaims.Subject

	subParts := strings.Split(sub, "|")
	if len(subParts) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Unauthorized."}`))
		return
	}

	userID := subParts[1]

	// Can use this userID for associations with this specific user
	fmt.Println(userID)

	responseData := map[string]string{"message": "Welcome to the private page"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
