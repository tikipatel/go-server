package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore collects data about players in memory
type InMemoryPlayerStore struct {
	scores map[string]int
}

// NewInMemoryPlayerStore returns an instance of InMemoryPlayerStore
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// GetPlayerScore retrieves scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

// GetLeague returns a slice of players
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

// RecordWin is a function that records win in memory
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}

func main() {
	inMemoryPlayerStore := NewInMemoryPlayerStore()
	server := NewPlayerServer(inMemoryPlayerStore)

	if err := http.ListenAndServe("localhost:5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
