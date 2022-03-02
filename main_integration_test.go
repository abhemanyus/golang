package main

import (
	"net/http"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	scores := Scores{
		"Pepper": 0,
	}
	store := CreatePlayerStore(scores)
	server := CreatePlayerServer(store)
	request, response := getFakeStuff("Pepper", http.MethodPost)
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)
	server.ServeHTTP(response, request)

	request, response = getFakeStuff("Pepper", http.MethodGet)
	server.ServeHTTP(response, request)

	assertStatus(t, http.StatusOK, response.Code)
	assertScore(t, "3", response.Body.String())
}
