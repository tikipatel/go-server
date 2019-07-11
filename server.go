package main

import (
	"fmt"
	"net/http"
)

// PlayerServer is a server that responds to HTTP requests
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
