package handler

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module exports all handler resources to Fx at startup
var Module = fx.Provide(
	New,
)

// Handler handles server requests
type Handler interface {
	Root(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger *zap.SugaredLogger
}

// New creates a new Handler
func New(
	logger *zap.SugaredLogger,
) (Handler, error) {
	return &handler{
		logger: logger,
	}, nil
}

// Root handles the root route
func (h *handler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
