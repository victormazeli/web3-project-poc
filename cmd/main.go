package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"goApiStartetProject/api/server"
	"goApiStartetProject/config"
	"log"
)

func main() {
	// Load config
	app := config.LoadEnvironmentVariables()
	env := app.Env

	// Connect to the database
	db, err := sqlx.Open("postgres", "host=127.0.0.1 user=cipher dbname=etheruem_poc port=5432 sslmode=disable password=cipher")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return
	}

	fmt.Println("Connected to the database!")

	// Initialize Server
	server.Server(env, db)
}
