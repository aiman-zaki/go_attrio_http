package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aiman-zaki/go_attrio_http/models"
	"github.com/aiman-zaki/go_attrio_http/services"
	"github.com/aiman-zaki/go_attrio_http/wrappers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type ContactUsResources struct {
}

func (rs ContactUsResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/", rs.Create)
		r.Get("/", rs.All)
	})

	return r
}

func (rs ContactUsResources) Create(w http.ResponseWriter, r *http.Request) {
	var m models.ContactUs
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	wrappers.JSONDecodeWrapper(w, r, &m)

	err := db.Insert(&m)
	if err != nil {
		fmt.Println(err)
	}
}

func (rs ContactUsResources) All(w http.ResponseWriter, r *http.Request) {
	var m []models.ContactUs
	db := pg.Connect(services.PgOptions())
	defer db.Close()
	err := db.Model(&m).Select()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(m)
}
