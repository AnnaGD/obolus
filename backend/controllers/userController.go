// controllers/userController.go

package controllers

import (
	"database/sql"
	"net/http"
	"encoding/json"
)

type User struct {

	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`

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

