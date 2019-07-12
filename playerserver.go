package main

import "net/http"

// PlayerStore describes an interface to provide players' scores.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// PlayerServer is a server for players
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
