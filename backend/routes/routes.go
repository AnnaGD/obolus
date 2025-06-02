// routes/routes.go

package routes

import (
	"obolus/backend/controllers"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Initialize ther route mapping
func InitRoutes(db *sql.DB) *chi.Mux{
	router := chi.NewRouter()

	// Set up CPRS middleware options
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins can be set to the origins that you expect to make requests from, or to "*" to allow any origin
		AllowedOrigins: []string{"*"}, // Adjust the port and domain as needed "http://localhost:3000", "https://yourdomain.com"

		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"}, // "Accept", "Authorization", "Content-Type", "X-CSRF-Token"

		ExposedHeaders: []string{},
		AllowCredentials: true,
		MaxAge: 200, // Maximum value not to be exceeded for Access-Control-Max-Age in seconds

	})

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// User routes
	router.Post("/signup", controllers.SignUpHandler(db))

	// Additional routes go here

	return router
}