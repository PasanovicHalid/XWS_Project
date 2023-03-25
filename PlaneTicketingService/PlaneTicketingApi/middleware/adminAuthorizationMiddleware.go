package middleware

import (
	"net/http"
	"planeTicketing/constants"
	"planeTicketing/services"
)

func MiddlewareAdminAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		jwt_claims := h.Context().Value(JwtContent{}).(*services.SignedDetails)

		if jwt_claims.Role != constants.AdminRole {
			http.Error(rw, "You are not an admin", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(rw, h)
	})
}
