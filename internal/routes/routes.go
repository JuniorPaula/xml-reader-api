package routes

import (
	"database/sql"
	"xml-reader-api/internal/config"
	"xml-reader-api/internal/handlers"
	authMiddleware "xml-reader-api/internal/middleware"
	"xml-reader-api/internal/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

const FILEPATH = "./data/Reconfile_fornecedores.xlsx"
const SHEETNAME = "Planilha1"

func NewRoutes(db *sql.DB, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	r.Use(middleware.WithValue("jwt-exp", cfg.JWTExp))

	userDB := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDB)
	authHandler := handlers.NewAuthHandler(userDB)

	supplierDB := repository.NewSupplierRepository(FILEPATH, SHEETNAME)
	supplierHandler := handlers.NewSupplierHandler(supplierDB)

	r.Post("/signup", userHandler.CreateUserHandler)
	r.Post("/login", authHandler.LoginHandler)

	r.Route("/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(authMiddleware.AuthenticatorMiddleware)

		r.Get("/suppliers", supplierHandler.GetSuppliersHandler)
	})

	return r
}
