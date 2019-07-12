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
	fmt.Println("Got here!! SERVEHTTP in server")
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
