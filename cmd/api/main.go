package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

type Application struct {
	Domain string
}

func main() {
	var app Application

	app.Domain = "example.com"

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.Routes(),
	}
	fmt.Println("server is listening to port", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
