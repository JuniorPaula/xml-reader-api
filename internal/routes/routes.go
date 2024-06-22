package routes

import (
	"database/sql"
	"xml-reader-api/internal/handlers"
	"xml-reader-api/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userDB := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/signup", userHandler.CreateUserHandler)

	})

	return r
}
