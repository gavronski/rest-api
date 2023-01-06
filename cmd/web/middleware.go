package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Auth checks if auth is correct
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login, password, ok := r.BasicAuth()

		if !ok {
			unauthorizedResponse(w)
			return
		}

		err := mainRepo.DB.Authenticate(login, password)

		if err != nil {
			unauthorizedResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

//
func unauthorizedResponse(w http.ResponseWriter) {
	jsonResponse := &jsonResponse{
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized",
	}

	out, _ := json.MarshalIndent(jsonResponse, "", "    ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(out)
}
