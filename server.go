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
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(player string) (score string) {
	if player == "Pepper" {
		score = "20"
	}
	if player == "Floyd" {
		score = "10"
	}
	return
}
