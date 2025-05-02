package controllers

import "net/http"

// HomeHandler responds to the "/" route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Obolus API"))
}