package configurations

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Server struct {
	config *Configurations
	mux    *runtime.ServeMux
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	return server
}

func (server *Server) Start() {
	log.Printf("Starting server on port %s", server.config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
