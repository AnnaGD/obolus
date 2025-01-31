package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Controller
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`,
	)
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handler)

	fmt.Println("Go backend started!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
