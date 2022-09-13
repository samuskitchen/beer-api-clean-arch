package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CORSMiddleware it is the function for the cors filter of the application
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Basic CORS Support
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"localhost"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})

		if r.Method == "OPTIONS" {
			_ = HTTPError(w, r, http.StatusNoContent, "")
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
