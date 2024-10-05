package main

import "net/http"

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
