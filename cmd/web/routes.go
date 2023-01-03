package main

import (
	"app/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/players", handlers.Repo.GetPlayers)
	mux.Get("/players/{id}", handlers.Repo.GetPlayer)
	mux.Post("/players", handlers.Repo.PostPlayer)
	mux.Patch("/players/{id}", handlers.Repo.UpdatePlayer)
	mux.Delete("/players/{id}", handlers.Repo.DeletePlayer)

	return mux
}
