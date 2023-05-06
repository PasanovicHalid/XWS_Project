package configurations

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/domain"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/infrastructure/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/persistance"
	mw "github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/startup/middlewares"
	authenticatePB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/authentification_service"
	userPB "github.com/PasanovicHalid/XWS_Project/BookingService/SharedLibraries/gRPC/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config     *Configurations
	mux        *runtime.ServeMux
	final_mux  *http.ServeMux
	keyService *application.KeyService
}

func NewServer(config *Configurations) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}

	mongo, err := persistance.NewMongoClient(config.ApiGatewayDbHost, config.ApiGatewayDbPort)
	if err != nil {
		log.Fatal(err)
	}

	server.initHandlers()

	keyRepository := persistance.NewKeyRepository(mongo)
	server.keyService = application.NewKeyService(keyRepository)
	jwtService := authentification.NewJwtService()

	//For scecific routes you have to build all your middlewares from scratch
	final_mux := http.NewServeMux()
	final_mux.Handle("/", mw.MiddlewareContentTypeSet(mw.MiddlewareAuthentification(server.mux, jwtService, server.keyService)))
	final_mux.Handle("/api/authenticate/login", mw.MiddlewareContentTypeSet(server.mux))
	final_mux.Handle("/api/authenticate/register", mw.MiddlewareContentTypeSet(server.mux))
	final_mux.Handle("/api/authenticate/getPublicKey", mw.MiddlewareContentTypeSet(server.mux))
	final_mux.Handle("/api/user/updateUser", mw.MiddlewareContentTypeSet(mw.MiddlewareAuthentification(mw.MiddlewareCheckIfUserRequestUsesIdentityOfLoggedInUser(server.mux, "identityId"), jwtService, server.keyService)))
	final_mux.Handle("/api/user/createUser", mw.MiddlewareContentTypeSet(mw.MiddlewareAuthentification(mw.MiddlewareCheckIfUserRequestUsesIdentityOfLoggedInUser(server.mux, "identityId"), jwtService, server.keyService)))
	final_mux.Handle("/getPublicKey", mw.MiddlewareContentTypeSet(server.GetPublicKeyHttp()))

	server.final_mux = final_mux

	return server
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
	go func() {
		server.GetPublicKeyForJwt()
	}()

	log.Printf("Starting server on port %s", server.config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.final_mux))
}

func (server *Server) GetPublicKeyForJwt() {
	time.Sleep(5 * time.Second)

	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/api/authenticate/getPublicKey", server.config.Port))

	if err != nil {
		log.Panic(err)
		return
	}

	if resp.StatusCode != 200 {
		log.Panic("Failed to get public key from authentification service")
		return
	}

	fields := make(map[string]string)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&fields)
	resp.Body.Close()

	server.keyService.SaveKey(&domain.Key{
		PublicKey: fields["publicKey"],
	})
}

func (server *Server) GetPublicKeyHttp() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		server.GetPublicKeyForJwt()

		rw.WriteHeader(http.StatusOK)
	})
}