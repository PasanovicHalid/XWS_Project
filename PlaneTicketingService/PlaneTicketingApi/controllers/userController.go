package controllers

import (
	"context"
	"net/http"
	"planeTicketing/constants"
	"planeTicketing/contracts"
	"planeTicketing/database"
	"planeTicketing/model"
	"planeTicketing/services"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

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

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
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

	token, err := services.GenerateToken(foundIdentity.Identity.Username, foundIdentity.Firstname, foundIdentity.Lastname, foundIdentity.Identity.Role, foundIdentity.Identity.Id.String())

	if err != nil {
		http.Error(rw, "Something failed when generating token", http.StatusInternalServerError)
		UserController.UserCollection.Logger.Panic(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(token))
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
