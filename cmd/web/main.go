package main

import (
	"app/internal/driver"
	"app/internal/handlers"
	"app/internal/models"
	"app/internal/repository"
	"app/internal/repository/dbrepo"
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var port string = ":8080"

// Repository is the repository type
type Repository struct {
	DB repository.DatabaseRepo
}

// Repo the repository used by the middleware
var mainRepo *Repository

// NewRepo creates new repository
func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

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

	dsn := os.Getenv("DSN")

	// connect to database
	log.Println("Connecting to databse ...")

	db, err := driver.ConnectSQL(dsn)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database!")

	mainRepo = NewRepo(db)
	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)
	return db, nil
}
