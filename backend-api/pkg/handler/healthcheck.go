package handler

import (
	"net/http"

	"github.com/unchainio/interfaces/logger"
	"gitlab.com/henkvanramshorst/forest/backend-api/pkg/backend-api"
)

type HealthCheckHandler struct {
	logger logger.Logger
	cfg    *backend_api.Config
}

func NewHealthCheckHandler(logger logger.Logger, cfg *backend_api.Config) *HealthCheckHandler {
	return &HealthCheckHandler{logger: logger, cfg: cfg}
}

func (h *HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`FOREST backend api`))
}
