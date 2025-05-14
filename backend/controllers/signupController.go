// controllers/signupController.go

package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {

	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"` // TODO: aka issue - in production store the password securely

}

// Checks the user fields for validity
func validateInput(user User) error {
	if user.Username == "" || len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if user.Password == "" || len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Regex for validating an email address
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if user.Email == "" || !emailRegex.MatchString(user.Email) {
		return errors.New("invalid e-mail format")
	}
	return nil
}

// New user registration
func SignUpHandler(db *sql.DB) http.HandlerFunc {

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

		// Validate input data
		if err := validateInput(newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Hash the pw before storing it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error while hashing the password.", http.StatusInternalServerError)
			return
		}

		// Check if the email already exits
		var existingId int
		err = db.QueryRow("SELECT id FROM users WHERE email = ?", newUser.Email).Scan(&existingId)
		if err != sql.ErrNoRows {
			if err != nil {
				http.Error(w, "Error checking for existing email.", http.StatusInternalServerError)
				return
			}
			// If we get here, it means that a user with this email already exists
			http.Error(w, "Email already in use.", http.StatusConflict)
			return
		}

		// Prepare the SQL statement using placeholders for parameters
		stmt, err := db.Prepare("INSERT INTO users(username, email, password_hash) VALUES(?, ?, ?)")
		if err != nil {
			http.Error(w, "Error preparing the SQL statement.", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(newUser.Username, newUser.Email, hashedPassword)
		if err != nil {
			// Handle the error properly. It could be due to a duplicate username or email, etc.
			http.Error(w, "Error creating the user account.", http.StatusInternalServerError)
			return
		}
		
		// For now pretend we've successfully created a user and return it
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"start": "user created"})
	}

}

