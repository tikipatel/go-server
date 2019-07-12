package main

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
