package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type Application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	var app Application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "postgres connection string")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	app.Domain = "example.com"

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.Routes(),
	}
	fmt.Println("server is listening to port", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
