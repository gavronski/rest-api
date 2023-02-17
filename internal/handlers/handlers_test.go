package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testData []struct {
	name           string
	url            string
	method         string
	postData       []byte
	expectedStatus int
}

func TestRepository_GetPlayers(t *testing.T) {
	testData := testData{
		{"get-players", "/players", "GET", nil, http.StatusOK},
	}

	for _, e := range testData {
		req := httptest.NewRequest(e.method, e.url, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.GetPlayers)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("Get players handler returns wrong response code %d", rr.Code)
		}
	}
}

func TestRepository_PostPlayer(t *testing.T) {
	testData := testData{
		{"post-players", "/players", "POST",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusOK},
		{"post-players", "/players", "POST",
			nil,
			http.StatusBadRequest},
		{"post-players", "/players", "POST",
			[]byte(
				`{"firstName":"Lionel","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusBadRequest},
	}

	for _, e := range testData {
		req := httptest.NewRequest(e.method, e.url, bytes.NewBuffer(e.postData))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.PostPlayer)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("Post players handler returns wrong response code %d", rr.Code)
		}
	}
}

func TestRepository_UpdatePlayer(t *testing.T) {
	testData := testData{
		{"patch-players", "/players/9", "PATCH",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusBadRequest},
		{"patch-players", "/players/1", "PATCH",
			nil,
			http.StatusBadRequest},
		{"patch-players", "/players/1", "PATCH",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusOK},
		{"patch-players", "/players", "PATCH",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusBadRequest},
	}

	for _, e := range testData {
		req := httptest.NewRequest(e.method, e.url, bytes.NewBuffer(e.postData))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.UpdatePlayer)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("Update players handler returns wrong response code %d", rr.Code)
		}
	}
}

func TestRepository_DeletePlayer(t *testing.T) {
	testData := testData{
		{"delete-player", "/players/1", "DELETE", nil, http.StatusOK},
		{"delete-player", "/players/9", "DELETE", nil, http.StatusBadRequest},
		{"delete-player", "/players", "DELETE", nil, http.StatusBadRequest},
	}

	for _, e := range testData {
		req := httptest.NewRequest(e.method, e.url, bytes.NewBuffer(e.postData))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.DeletePlayer)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("Delete player handler returns wrong response code %d", rr.Code)
		}
	}
}

func TestRepository_GetPlayer(t *testing.T) {
	testData := testData{
		{"delete-player", "/players/1", "GET",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusOK},
		{"delete-player", "/players/9", "GET",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusBadRequest},
		{"delete-player", "/players", "GET",
			[]byte(
				`{"firstName":"Cristiano","lastName":"Ronaldo","age":37,"country":"Portugal","club":"Manchester United","position":"striker","goals":24,"assists":6}`),
			http.StatusBadRequest},
	}

	for _, e := range testData {

		req := httptest.NewRequest(e.method, e.url, bytes.NewBuffer(e.postData))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.GetPlayer)
		handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatus {
			t.Errorf("Get player handler returns wrong response code %d", rr.Code)
		}
	}
}
