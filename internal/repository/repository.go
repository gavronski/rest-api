package repository

import "app/internal/models"

type DatabaseRepo interface {
	GetPlayers() ([]models.Player, error)
	InsertPlayer(player models.Player) error
}
