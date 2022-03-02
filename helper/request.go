package helper

import (
	"encoding/json"
	"net/http"
)

func CreateJSONDecoder(r *http.Request) *json.Decoder {
	return json.NewDecoder(r.Body)
}
