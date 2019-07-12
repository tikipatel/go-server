package main

import (
	"fmt"
	"log"
	"net/http"
)

// InMemoryPlayerStore collects data about players in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore retrieves scores for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

// RecordWin is a function that records win in memory
func (i *InMemoryPlayerStore) RecordWin(name string) {
	fmt.Printf("Recording win for player: %s\n", name)
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
