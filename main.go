package main

import (
	"log"
	"net/http"

	"github.com/aiman-zaki/go_attrio_http/handlers"
	"github.com/aiman-zaki/go_attrio_http/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {

	models.InitAttrio()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	r.Mount("/contact-us", handlers.ContactUsResources.Routes(handlers.ContactUsResources{}))
	log.Fatal(http.ListenAndServe(":3000", c.Handler(r)))
}
