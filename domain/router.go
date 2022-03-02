package domain

import (
	"log"
	"net/http"

	"github.com/andfxx27/chirps-api/connection"
	"github.com/andfxx27/chirps-api/domain/user"
	"github.com/andfxx27/chirps-api/middleware"
	"github.com/gorilla/mux"
)

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world from chirps api index handler!"))
}

func NewRouter() *mux.Router {
	db := connection.InitializeDatabaseConnection()
	if db != nil {
		log.Println("Connected to database!")
	}

	// Initialize repository and handler
	userRepository := user.NewRepository(db)

	userHandler := user.NewHandler(userRepository)

	router := mux.NewRouter()
	router.Use(middleware.JSONResponseMiddleware)
	router.HandleFunc("", indexHandler).Methods(http.MethodGet)

	userSubrouter := router.PathPrefix("/api/users").Subrouter()
	userSubrouter.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)
	userSubrouter.HandleFunc("/register", userHandler.Register).Methods(http.MethodPost)

	log.Println("Server up and running.")

	return router
}
