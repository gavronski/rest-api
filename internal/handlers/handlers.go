package handlers

import (
	"app/internal/driver"
	"app/internal/repository"
	"app/internal/repository/dbrepo"
	"encoding/json"
	"log"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	DB repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates new repository
func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewPostgresRepo(db.SQL),
	}
}

// Sets the repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := m.DB.GetPlayers()

	if err != nil {
		log.Println(err)
	}

	out, _ := json.MarshalIndent(players, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
