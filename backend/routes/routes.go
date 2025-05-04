// routes/routes.go

package routes

import (
	"obolus/backend/controllers"
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Initialize ther route mapping
func InitRoutes(db *sql.DB) *chi.Mux{
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", controllers.HomeHandler)

	// User routes
	router.Post("/signup", controllers.SignUpHandler(db))

	// Additional routes go here

	return router
}