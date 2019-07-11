package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		got := response.Body.String()
		want := "20"

		PlayerServer(response, request)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
