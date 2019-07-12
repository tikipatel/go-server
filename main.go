package main

import (
	"log"
	"net/http"
)

// InMemoryPlayerStore is a struct that will be used to hold on to player scores in memory
type InMemoryPlayerStore struct{}

// GetPlayerScore is a function that gets a player's score
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
