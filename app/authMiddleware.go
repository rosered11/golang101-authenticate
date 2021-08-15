package app

import (
	"net/http"

	"github.com/rosered11/golang101-lib/errors"

	"github.com/gorilla/mux"
	"github.com/rosered11/golang101-authenticate/domain"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVar := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVar)

				if isAuthorized {
					next.ServeHTTP(rw, r)
				} else {
					appError := errors.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					writeResponse(rw, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(rw, http.StatusUnauthorized, "missing token")
			}

		})
	}
}

func getTokenFromHeader(header string) string {
	return ""
}
