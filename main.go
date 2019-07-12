package main

import (
	"log"
	"net/http"
)

func main() {
	inMemoryPlayerStore := NewInMemoryPlayerStore()
	server := NewPlayerServer(inMemoryPlayerStore)

	if err := http.ListenAndServe("localhost:5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
