package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/andfxx27/chirps-api/constant"
	"github.com/andfxx27/chirps-api/helper"
	"github.com/golang-jwt/jwt"
)

func JSONResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			helper.BuildJSONResponse(http.StatusUnauthorized, constant.Unauthorized, "", rw)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			claimsJSON, _ := json.Marshal(claims)
			r.Header.Add("claims", string(claimsJSON))
		} else {
			helper.BuildJSONResponse(http.StatusUnauthorized, constant.Unauthorized, "", rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
