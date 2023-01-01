package main

import (
	"app/internal/driver"
	"app/internal/handlers"
	"app/internal/models"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var port string = ":8080"

func main() {
	db, err := run()

	if err != nil {
		log.Fatal(err)
	}

	// close db connection when app stops
	defer db.SQL.Close()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
func run() (*driver.DB, error) {
	gob.Register(models.Player{})

	// load all .env args
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbSSL := os.Getenv("DB_SSL")

	// Connect to database
	log.Println("Connecting to databse ...")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", dbHost, dbPort, dbName, dbUser, dbPass, dbSSL)
	db, err := driver.ConnectSQL(connectionString)

	if err != nil {
		log.Fatal(err)
	}

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)
	return db, nil
}
