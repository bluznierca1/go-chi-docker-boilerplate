package custommiddleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// AttachRequestIdHeader extracts already attached RequestID to context and adds new Header for Response with it
// Original chi.middleware does not do that
//
// Use it after using middleware.RequestID
func AttachRequestIdHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		receivedRequestId := middleware.GetReqID(r.Context())
		if receivedRequestId != "" {
			w.Header().Set(middleware.RequestIDHeader, receivedRequestId)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
