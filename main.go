package main

import (
	"log"
	"net/http"
	"os"
	"rest-and-go/store"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := store.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedOrigins := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, AllowedMethods)(router)))
}
