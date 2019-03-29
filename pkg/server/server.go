package server

import (
	"context"
	"log"
	"net/http"
)

// Server is interface defining methods required of the server
type Server interface {
	Start() error
	Stop() error
}

// Srv is an implementation of the Server interface
type Srv struct {
	port   string
	server *http.Server
}

// NewHTTPServer returns a Srv instance
func NewHTTPServer(port string) *Srv {
	return &Srv{
		port:   port,
		server: &http.Server{Addr: ":" + port},
	}
}

// Start causes the server to listen for tcp connections on specified port
func (s *Srv) Start() error {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server unexpectedly shutdown: %s", err)
	}

	return nil
}

// Stop causes the server to stop listening for tcp connections
func (s *Srv) Stop() error {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		return err
	}
	return nil
}
