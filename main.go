package main

import (
	"log"
	"net/http"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, ctx *Context) {
	w.Write([]byte(ctx.User.Name))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", NewRouteHandler(HomeHandler))

	s := &http.Server{
		Addr:           ":8000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
