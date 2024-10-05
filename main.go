package main

import (
	"log"
	"net/http"
	"time"
)

type HomeHandler struct {
	Context *Context
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Context.User.Name))
}

func NewHomeHandler() *MiddlewareHandler {
	h := &HomeHandler{
		&Context{},
	}
	return &MiddlewareHandler{
		Context: h.Context,
		Handler: h,
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", NewHomeHandler())

	s := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
