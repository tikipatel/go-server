package main

import (
	"fmt"
	"net/http"
)

// PlayerStore describes an interface to provide players' scores.
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a server for players
type PlayerServer struct {
	store PlayerStore
}

// ServeHTTP is a function to serve HTTP content
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	playerName := r.URL.Path[len("/players/"):]
	score := p.store.GetPlayerScore(playerName)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}
