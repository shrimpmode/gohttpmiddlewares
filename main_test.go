package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupHandler(h *http.Handler) {
}

func TestHome(t *testing.T) {
	tc := struct {
		want string
	}{
		want: "Giorno Giovanna",
	}

	h := NewHomeHandler()
	rr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	h.ServeHTTP(rr, r)

	got := rr.Body.String()
	if got != tc.want {
		t.Errorf("Got response: %q. Want: %q", got, tc.want)
	}
}
