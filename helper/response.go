package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func BuildJSONResponse(status int, message, errorLog string, rw http.ResponseWriter) {
	if errorLog != "" {
		log.Println(errorLog)
	}

	rw.WriteHeader(status)

	mapResponse := map[string]interface{}{
		"message": message,
	}
	jsonResponse, _ := json.Marshal(mapResponse)

	rw.Write(jsonResponse)
}

func BuildJSONResponseWithResult(status int, message string, result interface{}, errorLog string, rw http.ResponseWriter) {
	if errorLog != "" {
		log.Println(errorLog)
	}

	rw.WriteHeader(status)

	mapResponse := map[string]interface{}{
		"message": message,
		"result":  result,
	}
	jsonResponse, _ := json.Marshal(mapResponse)

	rw.Write(jsonResponse)
}
