package main

import (
	"encoding/json"
	"io"
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
func (i *InMemoryPlayerStore) GetLeague() League {
	var league League
	for name, wins := range i.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

// RecordWin is a function that records win in memory
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}

// FileSystemPlayerStore struct
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague implementation for file system player store
func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore from file system player store
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin is a function
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
		}
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}

func main() {
	inMemoryPlayerStore := NewInMemoryPlayerStore()
	server := NewPlayerServer(inMemoryPlayerStore)

	if err := http.ListenAndServe("localhost:5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
