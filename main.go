package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andfxx27/chirps-api/domain"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("An error occurred. Err: " + err.Error())
	}

	applicationName := os.Getenv("APPLICATION_NAME")
	applicationPort := os.Getenv("APPLICATION_PORT")
	log.Println(fmt.Sprintf("Starting %v on port %v...", applicationName, applicationPort))
	log.Fatalln(http.ListenAndServe("localhost:"+applicationPort, domain.NewRouter()))
}
