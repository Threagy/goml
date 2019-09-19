package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

// DefaultDriver implements the driver.Server interface. The zero value is a valid http.Server.
type DefaultDriver struct {
	Server http.Server
}

// NewDefaultDriver creates a driver with an http.Server with default timeouts.
func NewDefaultDriver() *DefaultDriver {
	return &DefaultDriver{
		Server: http.Server{
			ReadTimeout:  120 * time.Second,
			WriteTimeout: 120 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

// ListenAndServe sets the address and handler on DefaultDriver's http.Server,
// then calls ListenAndServe on it.
func (dd *DefaultDriver) ListenAndServe(addr string, h http.Handler) error {
	dd.Server.Addr = addr
	dd.Server.Handler = h

	log.Printf("serving on %s", addr)

	return dd.Server.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any active connections,
// by calling Shutdown on DefaultDriver's http.Server
func (dd *DefaultDriver) Shutdown(ctx context.Context) error {
	return dd.Server.Shutdown(ctx)
}
