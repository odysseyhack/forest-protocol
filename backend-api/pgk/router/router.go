package router

import (
	"github.com/unchainio/interfaces/logger"
	"github.com/go-chi/chi"
	"github.com/goware/cors"
	"gitlab.com/henkvanramshorst/forest/backend-api/pgk/handler"
)

func NewRouter(logger logger.Logger) (*chi.Mux, error) {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	r := chi.NewRouter()
	r.Use(c.Handler)

	// Setup handlers
	healthCheckHandler := handler.NewHealthCheckHandler(logger)

	// Setup services

	// Setup routes

	// Health Check
	r.Get("/", healthCheckHandler.HealthCheck)

	return r, nil
}