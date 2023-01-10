package dbrepo

import (
	"app/internal/models"
	"errors"
	"log"
)

// GetPlayers - selects all player from players table
func (m *testDBRepo) GetPlayers() ([]models.Player, error) {
	var players = []models.Player{}

	return players, nil
}

// InsertPlayers - adds player record to players table
func (m *testDBRepo) InsertPlayer(player models.Player) error {
	if player.FirstName == "Lionel" {
		log.Println("lionel")
		return errors.New("error while inserting player")
	}

	return nil
}

// UpdatePlayer - updates player fields
func (m *testDBRepo) UpdatePlayer(player models.Player) error {
	if player.ID == 9 {
		return errors.New("error while inserting player")
	}

	return nil
}

// DeletePlayer - deletes player row from the table
func (m *testDBRepo) DeletePlayer(id int) error {
	if id == 9 {
		return errors.New("error while inserting player")
	}

	return nil
}

// GetPlayerByID - select player from players table
func (m *testDBRepo) GetPlayerByID(id int) (models.Player, error) {
	var player models.Player

	if id == 9 {
		return player, errors.New("error while inserting player")
	}
	return player, nil
}

// Authenticate - compare data from the request and db
func (m *testDBRepo) Authenticate(login, testPassword string) error {
	return nil
}

// GetMaxID retruns max id from players table
func (m *testDBRepo) GetMaxID() (int, error) {
	return 1, nil
}
