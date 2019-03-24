package server

import (
	"context"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	Stop() error
}

type Srv struct {
	port   string
	server *http.Server
}

func NewHttpServer(port string) *Srv {
	return &Srv{
		port:   port,
		server: &http.Server{Addr: ":" + port},
	}
}

func (s *Srv) Start() error {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Server unexpectedly shutdodn: %s", err)
	}

	return nil
}

func (s *Srv) Stop() error {
	if err := s.server.Shutdown(context.TODO()); err != nil {
		return err
	}
	return nil
}
