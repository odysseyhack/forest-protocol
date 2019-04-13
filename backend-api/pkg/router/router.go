package router

import (
	"github.com/go-chi/chi"
	"github.com/goware/cors"
	"github.com/unchainio/interfaces/logger"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/handler"
)

func NewRouter(logger logger.Logger, cfg *backend_api.Config) (*chi.Mux, error) {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	})

	r := chi.NewRouter()
	r.Use(c.Handler)

	// Setup handlers
	healthCheckHandler := handler.NewHealthCheckHandler(logger, cfg)
	blockchainHandler := handler.NewBlockchainHandler(logger, cfg)
	locationHandler := handler.NewLocationHandler(logger, cfg)

	// Setup services

	// Health Check
	r.Get("/", healthCheckHandler.HealthCheck)

	// FeatherAPI
	r.Get("/blockchain/campaigns", blockchainHandler.Campaigns)
	r.Get("/blockchain/conversations", blockchainHandler.Conversations)
	r.Get("/blockchain/dacs", blockchainHandler.DACs)
	r.Get("/blockchain/donations", blockchainHandler.Donations)
	r.Get("/blockchain/events", blockchainHandler.Events)
	r.Get("/blockchain/milestones", blockchainHandler.Milestones)
	r.Get("/blockchain/pledgeAdmins", blockchainHandler.PledgeAdmins)
	r.Get("/blockchain/users", blockchainHandler.Users)
	r.Get("/blockchain/whitelist", blockchainHandler.Whitelist)

	r.Post("/location", locationHandler.GetLocationData)

	return r, nil
}
