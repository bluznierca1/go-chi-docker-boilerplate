package custommiddleware

import (
	"net/http"
	"os"
)

// SecurityHeaders attaches baic security headers to Writer
//
// It does not include CORS policy headers
func SecurityHeaders(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		// HSTS is added only in production since dev is usually http
		if os.Getenv("APP_ENV") != "production" {
			w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}

		// Content-Type is added when we generate response, in case of critical error, it will be default "text/html"
		w.Header().Add("X-Content-Type-Options", "nosniff")

		w.Header().Add("X-Frame-Options", "DENY")

		w.Header().Add("X-XSS-Protection", "1; mode=block")

		// Some of data is quite sensitive, so no caching
		w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
