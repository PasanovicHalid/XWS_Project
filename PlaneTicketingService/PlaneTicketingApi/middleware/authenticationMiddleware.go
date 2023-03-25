package middleware

import (
	"context"
	"net/http"
	"planeTicketing/services"
	"strings"
)

type JwtContent struct{}

func MiddlewareAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_bearer := h.Header.Get("Authorization")

		parts := strings.Split(jwt_bearer, "Bearer ")

		if len(parts) < 2 {
			http.Error(rw, "Token not present", http.StatusBadRequest)
			return
		}

		jwt := parts[1]

		claims, result := services.ValidateToken(jwt)

		if result == services.Token_Expired {
			http.Error(rw, "Token expired", http.StatusBadRequest)
			return
		}

		if result == services.Token_Invalid {
			http.Error(rw, "Token invalid", http.StatusBadRequest)
			return
		}

		if result == services.Token_Failed {
			http.Error(rw, "Something failed with token", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(h.Context(), JwtContent{}, claims)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
