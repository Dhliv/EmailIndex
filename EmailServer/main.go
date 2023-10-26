package main

import (
	"log"
	"net/http"

	"github.com/dhliv/EmailServer/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://localhost:5173"},
    //AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"*"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))
	r.Get("/Emails/{page}", handlers.GetEmails)

	serverPort := ":4052"
	log.Printf("MamuroEmail running in http://localhost%v\n", serverPort)
	err := http.ListenAndServe(serverPort, r)
	if err != nil {
		log.Printf("Error while listening ans serving: %v\n", err)
	}
}