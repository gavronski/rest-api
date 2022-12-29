package dbrepo

import (
	"app/internal/models"
)

func (m *postgresDBRepo) GetPlayers() ([]models.Player, error) {
	var players = []models.Player{}
	return players, nil
}
