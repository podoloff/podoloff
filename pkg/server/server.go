package server

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type contextKey string

func (c contextKey) String() string {
	return "server context key: " + string(c)
}

// Server is interface defining methods required of the server
type Server interface {
	Start() error
	Stop() error
}

// Srv is an implementation of the Server interface
type Srv struct {
	port    string
	connStr string
	db      *mongo.Client
	server  *http.Server
	cache   map[string]string
}

// NewHTTPServer returns a Srv instance
func NewHTTPServer(port string, dbconn string) *Srv {
	return &Srv{
		port:    port,
		connStr: dbconn,
		server:  &http.Server{Addr: ":" + port},
		cache:   make(map[string]string),
	}
}

// Start causes the server to listen for tcp connections on specified port
func (s *Srv) Start() error {
	clientOptions := options.Client().ApplyURI(s.connStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to Mongo.")

	s.db = client

	s.server.Handler = registerEndpoints(s)

	if err = s.server.ListenAndServe(); err != http.ErrServerClosed {
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

func registerEndpoints(s *Srv) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		s.createUser(w, r)
	})

	mux.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		s.authUser(w, r)
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		s.authTest(w, r)
	})

	mux.HandleFunc("/registerorg", func(w http.ResponseWriter, r *http.Request) {
		s.createOrg(w, r)
	})

	mux.HandleFunc("/getorg", func(w http.ResponseWriter, r *http.Request) {
		s.getOrg(w, r)
	})

	return mux
}
