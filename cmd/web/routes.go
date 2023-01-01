package main

import (
	"app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/players", handlers.Repo.GetPlayers)
	return mux
}
