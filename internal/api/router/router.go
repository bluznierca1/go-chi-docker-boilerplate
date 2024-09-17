package router

import (
	custommiddleware "myapp/internal/api/custom_middleware"
	"myapp/internal/api/handler"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(apiHandlers handler.ApiHandlers) chi.Router {
	router := chi.NewRouter()

	// attach security headers
	router.Use(custommiddleware.SecurityHeaders)
	router.Use(middleware.RequestID)                   // assign RequestId to context
	router.Use(custommiddleware.AttachRequestIdHeader) // make sure to attach Request ID to header in response
	router.Use(middleware.RealIP)                      // attempt to extract real IP from request
	router.Use(middleware.Heartbeat("/healthy"))       // add endpoint for pings

	// add profiling endpoints when not production
	if os.Getenv("APP_ENV") != "production" {
		router.Mount("/debug", middleware.Profiler())
	}

	router.Route("/api", func(r chi.Router) {
		// Group and define your routes in here
		r.Get("/hello", apiHandlers.HelloHandler.Hello)
	})

	return router
}
