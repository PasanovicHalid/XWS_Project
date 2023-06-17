package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"planeTicketing/constants"
	"planeTicketing/contracts"
	"planeTicketing/database"
	"planeTicketing/middleware"
	"planeTicketing/model"
	"planeTicketing/services"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApiKeyContent struct{}

type UserControllerDependecies struct {
	UserCollection *database.DatabaseCollection
}

var UserController *UserControllerDependecies

func SetupUserControllerRoutes(router *mux.Router) {
	signUpAdminRouter := router.Methods(http.MethodPost).Subrouter()
	signUpAdminRouter.HandleFunc("/users/signup/admin", SignUpAdmin)
	signUpAdminRouter.Use(MiddlewareSignUpDeserialization)

	signUpCustomerRouter := router.Methods(http.MethodPost).Subrouter()
	signUpCustomerRouter.HandleFunc("/users/signup/customer", SignUpCustomer)
	signUpCustomerRouter.Use(MiddlewareSignUpDeserialization)

	loginRouter := router.Methods(http.MethodPost).Subrouter()
	loginRouter.HandleFunc("/users/login", Login)
	loginRouter.Use(MiddlewareLoginDeserialization)

	userInfoRouter := router.Methods(http.MethodGet).Subrouter()
	userInfoRouter.HandleFunc("/users/info", GetUserInfo)
	userInfoRouter.Use(middleware.MiddlewareAuthentication)
	userInfoRouter.Use(MiddlewareUserInfoDeserialization)

	apiKeyRouter := router.Methods(http.MethodPost).Subrouter()
	apiKeyRouter.HandleFunc("/users/generate-api-key", GenerateApiKey)
	apiKeyRouter.Use(middleware.MiddlewareAuthentication)
	apiKeyRouter.Use(MiddlewareApiKeyGenerationDeserialization)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	getRouter.Use(MiddlewareApiKeyAuthorization)
}

func SignUpCustomer(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	signUpContract := h.Context().Value(KeyProduct{}).(*contracts.SignUpContract)

	customer := setupCustomer(signUpContract)

	count, err := UserController.UserCollection.Collection.CountDocuments(ctx, bson.M{"username": customer.User.Identity.Username})
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while counting usernames", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	if count > 0 {
		http.Error(rw, "Username already exists", http.StatusBadRequest)
		return
	}

	customer.User.Identity.Password = services.HashPassword(customer.User.Identity.Password)

	result, err := UserController.UserCollection.Collection.InsertOne(ctx, customer)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while adding customer", http.StatusInternalServerError)
		return
	}

	UserController.UserCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	rw.WriteHeader(http.StatusCreated)
}

func SignUpAdmin(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	signUpContract := h.Context().Value(KeyProduct{}).(*contracts.SignUpContract)

	admin := setupAdmin(signUpContract)

	count, err := UserController.UserCollection.Collection.CountDocuments(ctx, bson.M{"username": admin.Identity.Username})
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while counting usernames", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	if count > 0 {
		http.Error(rw, "Username already exists", http.StatusBadRequest)
		return
	}

	admin.Identity.Password = services.HashPassword(admin.Identity.Password)

	result, err := UserController.UserCollection.Collection.InsertOne(ctx, admin)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed while adding admin", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	UserController.UserCollection.Logger.Printf("Documents ID: %v\n", result.InsertedID)
	rw.WriteHeader(http.StatusCreated)
}

func Login(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	loginContract := h.Context().Value(KeyProduct{}).(*contracts.LoginContract)
	var foundIdentity model.User

	err := UserController.UserCollection.Collection.FindOne(ctx, bson.M{"username": loginContract.Username}).Decode(&foundIdentity)
	defer cancel()

	if err != nil {
		http.Error(rw, "Login credentials are incorrect", http.StatusBadRequest)
		return
	}

	passwordIsValid := services.VerifyPassword(loginContract.Password, foundIdentity.Identity.Password)
	defer cancel()

	if !passwordIsValid {
		http.Error(rw, "Login credentials are incorrect", http.StatusBadRequest)
		return
	}

	token, err := services.GenerateToken(foundIdentity.Identity.Username, foundIdentity.Firstname, foundIdentity.Lastname, foundIdentity.Identity.Role, foundIdentity.Identity.Id.Hex())

	if err != nil {
		http.Error(rw, "Something failed when generating token", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(token))
}

func GetUserInfo(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	userInfoContract := h.Context().Value(KeyProduct{}).(*contracts.UserInfoContract)

	userId, _ := primitive.ObjectIDFromHex(userInfoContract.Id)

	var foundIdentity model.User

	err := UserController.UserCollection.Collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&foundIdentity)
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed when getting user info", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(foundIdentity)
}

func GenerateApiKey(rw http.ResponseWriter, h *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	apiKeyGeneration := h.Context().Value(KeyProduct{}).(*contracts.ApiKeyGenerationContract)

	apiKey := services.GenerateApiKey(128)
	var apiKeyDuration time.Time

	if apiKeyGeneration.DurationForever {
		apiKeyDuration = time.Now().Add(time.Hour * 24 * 100000)
	} else {
		var err error
		apiKeyDuration, err = time.Parse(time.RFC3339, apiKeyGeneration.Duration)

		if err != nil {
			http.Error(rw, "Something failed when generating Token", http.StatusInternalServerError)
			UserController.UserCollection.Logger.Panic(err)
			return
		}
	}

	userId, _ := primitive.ObjectIDFromHex(apiKeyGeneration.UserId)

	result, err := UserController.UserCollection.Collection.UpdateOne(ctx, bson.M{"_id": userId}, bson.M{"$set": bson.M{"apiKey": apiKey, "apiKeyDuration": apiKeyDuration}})
	defer cancel()

	if err != nil {
		http.Error(rw, "Something failed when generating Token", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(rw, "Something failed when generating Token", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic("No user found with this id")
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func MiddlewareUserInfoDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		userInfoContract := &contracts.UserInfoContract{}
		jwtContent := h.Context().Value(middleware.JwtContent{}).(*services.SignedDetails)

		userInfoContract.Id = jwtContent.Uid

		ctx := context.WithValue(h.Context(), KeyProduct{}, userInfoContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareApiKeyGenerationDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		apiKeyGenerationContract := &contracts.ApiKeyGenerationContract{}
		jwtContent := h.Context().Value(middleware.JwtContent{}).(*services.SignedDetails)

		err := apiKeyGenerationContract.FromJSON(h.Body)

		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			UserController.UserCollection.Logger.Panic(err)
			return
		}

		apiKeyGenerationContract.UserId = jwtContent.Uid

		ctx := context.WithValue(h.Context(), KeyProduct{}, apiKeyGenerationContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareApiKeyAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		api_bearer := h.Header.Get("APIKey")

		parts := strings.Split(api_bearer, "Bearer ")

		if len(parts) < 2 {
			http.Error(rw, "Token not present", http.StatusBadRequest)
			return
		}

		apiKey := parts[1]

		var foundIdentity model.User

		err := UserController.UserCollection.Collection.FindOne(h.Context(), bson.M{"apiKey": apiKey}).Decode(&foundIdentity)

		if err != nil {
			http.Error(rw, "Unauthorized", http.StatusInternalServerError)
			UserController.UserCollection.Logger.Panic(err)
			return
		}

		if foundIdentity.Identity.ApiKeyDuration.Before(time.Now().UTC()) {
			http.Error(rw, "Token Expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(h.Context(), ApiKeyContent{}, &foundIdentity)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareSignUpDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		signUpContract := &contracts.SignUpContract{}
		err := signUpContract.FromJSON(h.Body)

		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			UserController.UserCollection.Logger.Panic(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, signUpContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareLoginDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		loginContract := &contracts.LoginContract{}
		err := loginContract.FromJSON(h.Body)

		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			UserController.UserCollection.Logger.Panic(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, loginContract)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func setupCustomer(signUpContract *contracts.SignUpContract) *model.Customer {
	customer := &model.Customer{
		User: model.User{
			Identity: model.Identity{
				Role:     constants.CustomerRole,
				Username: signUpContract.Username,
				Password: signUpContract.Password,
			},
			Firstname: signUpContract.Firstname,
			Lastname:  signUpContract.Lastname,
		},
	}
	return customer
}

func setupAdmin(signUpContract *contracts.SignUpContract) *model.User {
	admin := &model.User{
		Identity: model.Identity{
			Role:     constants.AdminRole,
			Username: signUpContract.Username,
			Password: signUpContract.Password,
		},
		Firstname: signUpContract.Firstname,
		Lastname:  signUpContract.Lastname,
	}
	return admin
}
