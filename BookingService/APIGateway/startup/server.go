package configurations

import (
	"context"
	"fmt"
	"log"
	"net/http"

	authenticatePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config    *Configurations
	mux       *runtime.ServeMux
	final_mux *http.ServeMux
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	server.initHandlers()

	auth_mux := http.NewServeMux()
	auth_mux.Handle("/", MiddlewareAdminAuthorization1(server.mux))
	auth_mux.Handle("/api/authenticate/login", MiddlewareAdminAuthorization(server.mux))
	auth_mux.Handle("/api/authenticate/register", MiddlewareAdminAuthorization(server.mux))

	server.final_mux = auth_mux

	return server
}

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Test middleware")
		c.Next()
	}
}

func MiddlewareAdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println("Test middleware")

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareAdminAuthorization1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println("Test middleware1")

		next.ServeHTTP(rw, h)
	})
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	authentificationEndpoint := fmt.Sprintf("%s:%s", server.config.AuthentificationHost, server.config.AuthentificationPort)
	err := authenticatePB.RegisterAuthenticateServiceHandlerFromEndpoint(context.TODO(), server.mux, authentificationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = userPB.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	log.Printf("Starting server on port %s", server.config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.final_mux))
}
