package repository

import "app/internal/models"

type DatabaseRepo interface {
	GetPlayers() ([]models.Player, error)
	InsertPlayer(player models.Player) error
	UpdatePlayer(player models.Player) error
	DeletePlayer(id int) error
	GetPlayerByID(id int) (models.Player, error)
	Authenticate(login, testPassword string) error
	GetMaxID() (int, error)
}
