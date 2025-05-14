// controllers/signupController.go

package controllers

import (
	"database/sql"
	"net/http"
	"encoding/json"
	"errors"
)

type User struct {

	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"` // TODO: aka issue - in production store the password securely

}

// Checks the user fields for validity

func validateInput(user User) error {
	if user.Username == "" || len(user.Username) < 3 {
		return errors.New("Username must be at least 3 characters long")
	}
	return nil
}

// New user registration

func SignUpHandler(sb *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		// Access to db in the handler
		if r.Method != "POST" {
			http.Error (w, "Method is not supported", http.StatusMethodNotAllowed)
			return
		}
		
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		// For now pretend we've successfully created a user and return it
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}

}

