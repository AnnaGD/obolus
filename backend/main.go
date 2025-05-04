package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"obolus/backend/routes"
	"github.com/go-chi/docgen"
)

// passing the routes flag to print docs -- to run it: `go run . -routes`
var printRoutes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse() // see if the route was passed to generate documentation

	router := routes.InitRoutes() // create a new router

	fmt.Println("Go backend started!")

	// Passing -routes to the program will generate docs for the above router definition.
	// See the `routes.json` file in this folder for the output.

	if *printRoutes {
		fmt.Println(docgen.MarkdownRoutesDoc(router, docgen.MarkdownOpts{
			ProjectPath: "https://github.com/AnnaGD/obolus",
			Intro: "Welcome to Obolus, your fave travel coordinator!",
		}))
		return
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
