package middleware

import (
	"errors"
	"net/http"
	"xml-reader-api/internal/utils"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
)

var errUnauthorized = errors.New("Unauthorized")

// AuthenticatorMiddleware is a middleware that checks if the user is authenticated
func AuthenticatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			utils.ErrorJSON(w, errUnauthorized, http.StatusUnauthorized)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			utils.ErrorJSON(w, errUnauthorized, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
