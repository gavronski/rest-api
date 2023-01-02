package handlers

import (
	"app/internal/driver"
	"app/internal/models"
	"app/internal/repository"
	"app/internal/repository/dbrepo"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// GetPlayers - returns all players in JSON format
func (m *Repository) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := m.DB.GetPlayers()

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	out, _ := json.MarshalIndent(players, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// PostPlayer - decodes JSON to Player and inserts record to players table
func (m *Repository) PostPlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	// decode request's body
	err := json.NewDecoder(r.Body).Decode(&player)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// insert player into db
	err = m.DB.InsertPlayer(player)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, "ok")
}

func (m *Repository) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	id, err := getID(r.URL)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	player.ID = id

	// decode request's body
	err = json.NewDecoder(r.Body).Decode(&player)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	// update player fields
	err = m.DB.UpdatePlayer(player)

	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	responseJSON(w, http.StatusOK, "ok")
}

// responseJSON sends JSON response
func responseJSON(w http.ResponseWriter, status int, message string) {
	jsonResponse := &jsonResponse{
		Status:  status,
		Message: message,
	}

	out, _ := json.MarshalIndent(jsonResponse, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func getID(url *url.URL) (int, error) {
	// get only id from the address
	id := strings.Replace(url.Path, "/players/", "", -1)

	playerID, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return playerID, nil
}
