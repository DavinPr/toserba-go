package api

import (
	"context"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type HTTPAPIServer struct {
	server *http.Server
}

func NewHTTPServer(cfg ServerConfig, r http.Handler) *HTTPAPIServer {
	server := http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	return &HTTPAPIServer{
		server: &server,
	}
}

// Start starts the API server in a new goroutine.
func (h *HTTPAPIServer) Start() {
	func(h *HTTPAPIServer) {
		log.Info().Str("op", "ApiServer/Start").Int("PID", os.Getpid()).Msgf("Starting server on %s", h.server.Addr)

		err := h.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("Unhandled server shutdown")
		}
	}(h)
}

// Shutdown stops the HTTP server.
func (h *HTTPAPIServer) Shutdown() error {
	return h.server.Shutdown(context.Background())
}
