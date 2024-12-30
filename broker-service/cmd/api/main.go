package main

import (
	"fmt"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

const Port = "80"

type Config struct {
	Rabbit *amqp.Connection
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
