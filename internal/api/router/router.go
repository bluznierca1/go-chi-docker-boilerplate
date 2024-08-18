package router

import (
	custommiddleware "myapp/internal/api/custom_middleware"
	"myapp/internal/api/handler"

	"github.com/go-chi/chi/v5"
)

func InitRouter(apiHandlers handler.ApiHandlers) chi.Router {
	router := chi.NewRouter()

	// attach security headers
	router.Use(custommiddleware.SecurityHeaders)

	// Group and define your routes in here
	router.Get("/hello", apiHandlers.HelloHandler.Hello)

	return router
}
