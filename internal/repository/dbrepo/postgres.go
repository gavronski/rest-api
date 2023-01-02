package dbrepo

import (
	"app/internal/models"
	"context"
	"log"
	"time"
)

// GetPlayers - selects all player from players table
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

// InsertPlayers - adds player record to players table
func (m *postgresDBRepo) InsertPlayer(player models.Player) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// use nextval('players_sequence') to increment id's value
	stmt := `
	insert into players 
		(id, first_name, last_name, age, country, club, position, goals, assists, created_at, updated_at)
	values
		(nextval('players_sequence'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := m.DB.ExecContext(ctx, stmt,
		player.FirstName,
		player.LastName,
		player.Age,
		player.Country,
		player.Club,
		player.Position,
		player.Goals,
		player.Assists,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
