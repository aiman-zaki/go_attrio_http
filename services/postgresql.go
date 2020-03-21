package services

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

//PostGresDb to return db of connected postgres
func PgOptions() *pg.Options {
	return &pg.Options{

		User:     "postgres",
		Password: "",
		Database: "attrio",
	}

}

//CreateTable for postgres
func CreateTable(db *pg.DB, model interface{}) {
	err := db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists:   true,
		FKConstraints: true,
	})
	if err != nil {
		fmt.Println("Error during table creation", err)
	} else {
		fmt.Println("Table is created : ", model)
	}
}
