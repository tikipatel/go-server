package main

import (
	"fmt"
	"net/http"
)

// PlayerServer is a server that responds to HTTP requests
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	var score string
	if player == "Pepper" {
		score = "20"
	}
	if player == "Floyd" {
		score = "10"
	}

	fmt.Fprint(w, score)
}
