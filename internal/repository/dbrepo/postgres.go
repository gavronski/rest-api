package dbrepo

import (
	"app/internal/models"
	"context"
	"log"
	"time"
)

func (m *postgresDBRepo) GetPlayers() ([]models.Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var players = []models.Player{}

	query := `select * from players`
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return players, err
	}

	for rows.Next() {
		var player models.Player

		err = rows.Scan(
			&player.ID,
			&player.FirstName,
			&player.LastName,
			&player.Age,
			&player.Country,
			&player.Club,
			&player.Position,
			&player.Goals,
			&player.Assists,
			&player.CreatedAt,
			&player.UpdatedAt,
		)

		if err != nil {
			return players, err
		}
		log.Println(player)
		players = append(players, player)
	}
	return players, nil
}
