package helper

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func CreateJSONDecoder(r *http.Request) *json.Decoder {
	return json.NewDecoder(r.Body)
}

func ParseJWTClaims(r *http.Request) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	err := json.Unmarshal([]byte(r.Header.Get("claims")), &claims)
	return claims, err
}
