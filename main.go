package main

import (
	"log"
	"net/http"
	"time"
)

type Context struct {
	User *User
}

type AppHandler interface {
	SetUser(*User)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type MiddlewareHandler struct {
	Context *Context
	Handler http.Handler
}

func (f *MiddlewareHandler) SetUser(user *User) {
	f.Context.User = user
}

func (f *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := &AuthMiddleware{next: f}
	// Apply more middlewares here
	m.ServeHTTP(w, r)
}

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
