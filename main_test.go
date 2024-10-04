package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupHandler(h *http.Handler) {
}

func TestHome(t *testing.T) {
	h := NewHomeHandler()
	rr := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	h.ServeHTTP(rr, r)

	t.Errorf("Response: %q", rr.Body.String())
}
