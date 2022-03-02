package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLeague(t *testing.T) {
	scores := Scores{
		"Pepper": 42,
		"Floyd":  69,
	}
	store := CreatePlayerStore(scores)
	server := CreatePlayerServer(store)

	t.Run("return 200 on /league", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		gotHdr := response.Result().Header.Get("content-type")
		wantHdr := "application/json"
		if wantHdr != gotHdr {
			t.Errorf("want %q, got %q", wantHdr, gotHdr)
		}
		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Errorf("unable to parse response %q into slice of Players %v", response.Body, got)
		}
		assertStatus(t, http.StatusOK, response.Code)
	})
}

func TestGetPlayer(t *testing.T) {
	scores := Scores{
		"Pepper": 42,
		"Floyd":  69,
	}
	store := CreatePlayerStore(scores)
	server := CreatePlayerServer(store)
	t.Run("return Pepper's score", func(t *testing.T) {
		request, response := getFakeStuff("Pepper", http.MethodGet)

		server.ServeHTTP(response, request)
		assertStatus(t, http.StatusOK, response.Code)
		assertScore(t, "42", response.Body.String())
	})
	t.Run("return Floyd's score", func(t *testing.T) {
		request, response := getFakeStuff("Floyd", http.MethodGet)

		server.ServeHTTP(response, request)
		assertStatus(t, http.StatusOK, response.Code)
		assertScore(t, "69", response.Body.String())

	})
	t.Run("return 404 for missing player", func(t *testing.T) {
		request, response := getFakeStuff("Apollo", http.MethodGet)
		server.ServeHTTP(response, request)
		assertStatus(t, http.StatusNotFound, response.Code)
		assertScore(t, "player not found", response.Body.String())
	})
}

func TestStoreWin(t *testing.T) {
	scores := Scores{
		"Pepper": 42,
	}
	store := CreatePlayerStore(scores)
	server := CreatePlayerServer(store)
	t.Run("increase Pepper's score", func(t *testing.T) {
		request, response := getFakeStuff("Pepper", http.MethodPost)
		server.ServeHTTP(response, request)
		assertStatus(t, http.StatusAccepted, response.Code)
		if scores["Pepper"] != 43 {
			t.Errorf("want %v, got %v for player %q", 43, scores["Pepper"], "Pepper")
		}
	})
}

func getFakeStuff(name, method string) (*http.Request, *httptest.ResponseRecorder) {
	return httptest.NewRequest(method, "/player/"+name, nil), httptest.NewRecorder()
}

func assertScore(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func assertStatus(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
