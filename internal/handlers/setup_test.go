package handlers

import (
	"app/internal/driver"
	"app/internal/models"
	"app/internal/repository"
	"app/internal/repository/dbrepo"
	"encoding/gob"
	"os"
	"testing"
)

// Repository is the repository type
type TestRepository struct {
	DB repository.DatabaseRepo
}

// Repo the repository used by the middleware
var mainRepo *Repository

// NewRepo creates new repository
func NewTestRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

func TestMain(m *testing.M) {
	gob.Register(models.Player{})

	// use testing db repo to not act with real db
	mainRepo = NewTestingRepo()
	repo := NewTestingRepo()
	NewHandlers(repo)

	os.Exit(m.Run())
}
