package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var invalidTokenError = "invalid token"

type AuthProps struct{}

func AuthenticationMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := &response{
			Success: false,
		}

		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			response.Error = "no token provided"
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		splitedToken := strings.Split(accessToken, " ")
		if len(splitedToken) != 2 {
			response.Error = invalidTokenError
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		jwtToken := splitedToken[1]
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New(invalidTokenError)
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			response.Error = err.Error()
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), AuthProps{}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		response.Error = invalidTokenError
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	})
}
