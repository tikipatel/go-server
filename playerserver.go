package main

import (
	"net/http"
)

// PlayerStore describes an interface to provide players' scores.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// PlayerServer is a server for players
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// Player is an object describing a player
type Player struct {
	Name string
	Wins int
}

// ReaderSeeker interface
type ReaderSeeker interface {
	Reader
	Seeker
}

// Reader interface
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Seeker Interface
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
