package models

import "time"

type Player struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Age       int       `json:"age"`
	Country   string    `json:"country"`
	Club      string    `json:"club"`
	Position  string    `json:"position"`
	Goals     int       `json:"goals"`
	Assists   int       `json:"assists"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
