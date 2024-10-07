package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting server on port %s\n", Port)

	//define http server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}

	//start the server

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}
