package main

import (
	"net/http"

	"github.com/aiman-zaki/go_attrio_http/handlers"
	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/contact-us", handlers.ContactUsResources.Routes(handlers.ContactUsResources{}))
	http.ListenAndServe(":3000", r)
}
