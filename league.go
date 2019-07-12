package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a typealias to an array of players
type League []Player

// NewLeague returns a new array of players or an error
func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league: %v", err)
	}
	return league, err
}
