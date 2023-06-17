package middlewares

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/application"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/infrastructure/authentification"
	"github.com/PasanovicHalid/XWS_Project/BookingService/APIGateway/presentation/contracts"
	"golang.org/x/exp/slices"
)

type JwtContent struct{}
type RequestDecoded struct{}

func MiddlewareAuthentification(next http.Handler, jwtService *authentification.JwtService, keyService *application.KeyService) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_bearer := h.Header.Get("Authorization")

		parts := strings.Split(jwt_bearer, "Bearer ")

		if len(parts) < 2 {
			http.Error(rw, "Token not present", http.StatusUnauthorized)
			return
		}

		jwt := parts[1]

		key, err := keyService.GetKey()

		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		claims, result := jwtService.ValidateToken(jwt, key.PublicKey)

		if result == authentification.Token_Expired {
			http.Error(rw, "Token expired", http.StatusUnauthorized)
			return
		}

		if result == authentification.Token_Invalid {
			http.Error(rw, "Token invalid", http.StatusUnauthorized)
			return
		}

		if result == authentification.Token_Failed {
			http.Error(rw, "Something failed with token", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(h.Context(), JwtContent{}, claims)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareContentTypeSetWithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		log.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		// Add CORS headers
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Add Content-Type header
		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareAuthorization(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_claims := h.Context().Value(JwtContent{}).(*authentification.SignedDetails)

		if !slices.Contains(roles, jwt_claims.Role) {
			http.Error(rw, "You are not authorized", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareDecodeRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		request := &contracts.FilterAccomodationRequest{}

		err := json.NewDecoder(h.Body).Decode(request)

		if err != nil {
			http.Error(rw, "Something went wrong. Unable to decode json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(h.Context(), RequestDecoded{}, request)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}

func MiddlewareCheckIfUserRequestUsesIdentityOfLoggedInUser(next http.Handler, fieldName string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_claims := h.Context().Value(JwtContent{}).(*authentification.SignedDetails)

		fields := make(map[string]interface{})

		bodyBytes, _ := io.ReadAll(h.Body)
		err := json.Unmarshal(bodyBytes, &fields)
		h.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))
		h.Body.Close()

		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		if jwt_claims.Id != fields[fieldName] {
			http.Error(rw, "Use id of your profile", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, h)
	})
}

func MiddlewareAddIdentityIdToRequest(next http.Handler, fieldName string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_claims := h.Context().Value(JwtContent{}).(*authentification.SignedDetails)

		fields := make(map[string]interface{})

		bodyBytes, _ := io.ReadAll(h.Body)
		err := json.Unmarshal(bodyBytes, &fields)

		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		fields[fieldName] = jwt_claims.Id

		fieldsJson, err := json.Marshal(fields)

		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		h.Body = io.NopCloser(strings.NewReader(string(fieldsJson)))

		h.Body.Close()

		if jwt_claims.Id != fields[fieldName] {
			http.Error(rw, "Use id of your profile", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(rw, h)
	})
}
