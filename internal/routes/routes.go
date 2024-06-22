package routes

import (
	"database/sql"
	"xml-reader-api/internal/config"
	"xml-reader-api/internal/handlers"
	"xml-reader-api/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRoutes(db *sql.DB, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwt-exp", cfg.JWTExp))

	userDB := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB)
	authHandler := handlers.NewAuthHandler(userDB)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/signup", userHandler.CreateUserHandler)
		r.Post("/login", authHandler.LoginHandler)

		// Protected routes
	})

	return r
}
