package models

import (
	"github.com/aiman-zaki/go_attrio_http/services"
	"github.com/go-pg/pg/v9"
)

func InitAttrio() {
	db := pg.Connect(services.PgOptions())
	defer db.Close()

	models := []interface{}{
		(*ContactUs)(nil),
	}

	for _, model := range models {
		services.CreateTable(db, model)
	}
}
