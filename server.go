package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

type Player struct {
	Name  string
	Score int
}

type Scores map[string]int

func CreatePlayerServer(store *Store) *PlayerServer {
	if store == nil {
		store = CreatePlayerStore(nil)
	}
	router := http.NewServeMux()
	playerServer := &PlayerServer{store, router}
	router.Handle("/league", http.HandlerFunc(playerServer.leagueHandler))
	router.Handle("/player/", http.HandlerFunc(playerServer.playerHandler))
	return playerServer
}

func CreatePlayerStore(scores Scores) *Store {
	return &Store{scores: scores, mu: sync.Mutex{}}
}

type PlayerStore interface {
	GetPlayerScore(string) (int, error)
	RecordPlayerWin(string) error
	GetPlayerList() []Player
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) postScore(w http.ResponseWriter, player string) {
	err := p.store.RecordPlayerWin(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetPlayerList())
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/player/")
	switch r.Method {
	case http.MethodGet:
		p.getScore(w, player)
	case http.MethodPost:
		p.postScore(w, player)
	}
}

type Store struct {
	scores map[string]int
	mu     sync.Mutex
}

func (s *Store) GetPlayerScore(name string) (int, error) {
	score, ok := s.scores[name]
	if ok {
		return score, nil
	}
	return 0, fmt.Errorf("player not found")
}

func (s *Store) RecordPlayerWin(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	score, ok := s.scores[name]
	if ok {
		s.scores[name] = score + 1
		return nil
	}
	return fmt.Errorf("player not found")
}

func (s *Store) GetPlayerList() []Player {
	var List []Player
	for name, score := range s.scores {
		List = append(List, Player{name, score})
	}
	return List
}
