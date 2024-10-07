package main

import (
	"auth/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const Port = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

var counts int64

func main() {
	log.Println("Starting server on port " + Port)

	conn := connectDB()

	if conn == nil {
		log.Fatal("Error connecting to database")
	}

	//set up config

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", Port),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to make sure it's alive
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Error connecting to database:", err)
			counts++
		} else {
			log.Println("connected to database")
			return connection
		}
		if counts > 10 {
			log.Fatal("Too many errors connecting to database")
			return nil
		}
		log.Println("Retrying connection to database")
		time.Sleep(time.Second * 2)
		continue
	}
}
