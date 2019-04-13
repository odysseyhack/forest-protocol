package handler

import (
	"github.com/unchainio/interfaces/logger"
	"net/http"
)

type HealthCheckHandler struct {
	logger logger.Logger
}

func NewHealthCheckHandler(logger logger.Logger) *HealthCheckHandler {
	return &HealthCheckHandler{logger: logger}
}

func (p *HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`FOREST backend api`))
}
